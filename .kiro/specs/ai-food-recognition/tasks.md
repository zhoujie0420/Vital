# Implementation Plan: AI Food Recognition

## Overview

Implement AI-powered food recognition for the Vital Fitness diet tracking feature. The backend adds a DashScope (qwen-vl) integration behind a new `POST /api/v1/foods/recognize` endpoint. The frontend extends `diet/add.vue` with a camera button, image picker, and recognition result sheet. Implementation proceeds backend-first (config → client → service → handler → route), then frontend (API client → UI components → integration), with property-based tests validating correctness properties from the design.

## Tasks

- [x] 1. Extend backend config with DashScope settings
  - [x] 1.1 Add `DashScope` struct to `internal/config/config.go`
    - Add `DashScope` struct with `APIKey`, `Model`, `BaseURL`, `Timeout` fields (mapstructure tags: `API_KEY`, `MODEL`, `BASE_URL`, `TIMEOUT`)
    - Add `DashScope` field to the `Config` struct
    - Add viper defaults: `MODEL` = `"qwen-vl-plus"`, `BASE_URL` = `"https://dashscope.aliyuncs.com/compatible-mode/v1"`, `TIMEOUT` = `15`
    - Add environment variable overrides for `DASHSCOPE_API_KEY`, `DASHSCOPE_MODEL`, `DASHSCOPE_BASE_URL`, `DASHSCOPE_TIMEOUT`
    - _Requirements: 2.1, 4.1_
  - [x] 1.2 Add DASHSCOPE section to `configs/config.yaml` and `configs/config.yaml.example`
    - Add `DASHSCOPE` block with `API_KEY`, `MODEL`, `BASE_URL`, `TIMEOUT` keys
    - _Requirements: 2.1_
  - [x] 1.3 Add `DASHSCOPE_API_KEY` environment variable to `docker-compose.yml`
    - Add `DASHSCOPE_API_KEY` to the `api` service environment list
    - _Requirements: 2.1_

- [x] 2. Implement DashScope client
  - [x] 2.1 Create `internal/service/dashscope_client.go`
    - Define `DashScopeClient` struct with `apiKey`, `model`, `baseURL`, `httpClient` fields
    - Implement `NewDashScopeClient(cfg config.DashScope) *DashScopeClient` constructor with configurable timeout
    - Implement `RecognizeFood(ctx context.Context, imageBase64 string, mimeType string) (string, error)` method
    - Construct OpenAI-compatible chat completion request with system prompt (food recognition JSON prompt from design), user message with base64 image URL, `temperature: 0.1`, `max_tokens: 2048`
    - Handle HTTP errors: return specific errors for timeout (504), non-200 DashScope response (502), missing API key (500)
    - _Requirements: 2.1, 2.3, 6.1, 6.2_

- [x] 3. Implement food recognition service with parsing and validation
  - [x] 3.1 Create `internal/service/food_recognition_service.go`
    - Define `RecognizedFood` struct and `FoodRecognitionResponse` struct matching the design data models
    - Define `FoodRecognizeRequest` struct with `Image` field (binding:"required")
    - Implement `NewFoodRecognitionService(cfg config.DashScope) *FoodRecognitionService`
    - Implement `RecognizeFood(ctx context.Context, imageBase64 string, mimeType string) (*FoodRecognitionResponse, error)` that validates input, calls DashScope client, and parses response
    - Implement `parseModelResponse(raw string) ([]RecognizedFood, error)` — extract JSON from model output (handle markdown code block wrapping), unmarshal into `[]RecognizedFood`
    - Implement `validateImageFormat(base64Data string) (string, error)` — decode first bytes, check JPEG magic (`FF D8 FF`) or PNG magic (`89 50 4E 47`), return detected MIME type or error
    - Implement `validateImageSize(base64Data string) error` — check decoded size ≤ 5MB
    - _Requirements: 2.2, 2.4, 2.5, 4.4, 4.5, 4.6_

  - [ ]* 3.2 Write property test: model response parsing preserves all food data (Property 1)
    - **Property 1: Model response parsing preserves all food data**
    - Generate random valid JSON with `foods` array (0–N items, each with name, estimated_grams, calories_per_100g, protein_per_100g, carbs_per_100g, fat_per_100g, confidence)
    - Assert `parseModelResponse` returns slice of same length with matching field values
    - Minimum 100 iterations
    - **Validates: Requirements 2.2, 2.4**

  - [ ]* 3.3 Write property test: image format validation accepts only JPG and PNG (Property 4)
    - **Property 4: Image format validation accepts only JPG and PNG**
    - Generate random byte sequences; prepend JPEG magic (`FF D8 FF`) or PNG magic (`89 50 4E 47`) for valid cases
    - Assert validator accepts if and only if magic bytes match JPEG or PNG
    - Assert rejected inputs produce error with appropriate message
    - Minimum 100 iterations
    - **Validates: Requirements 4.4, 4.6**

  - [ ]* 3.4 Write unit tests for parsing edge cases
    - Test `parseModelResponse` with: valid single-food JSON, multi-food JSON, empty foods array, malformed JSON, JSON wrapped in markdown code blocks, extra text around JSON
    - Test `validateImageFormat` with JPEG, PNG, GIF, BMP, and random bytes
    - Test `validateImageSize` at boundary (4.99MB, 5MB, 5.01MB)
    - _Requirements: 2.2, 2.5, 4.4, 4.5_

- [x] 4. Checkpoint — Ensure all backend tests pass
  - Ensure all tests pass, ask the user if questions arise.

- [x] 5. Implement food recognition HTTP handler and register route
  - [x] 5.1 Create `internal/handler/food_recognition_handler.go`
    - Define `FoodRecognitionHandler` struct with `service *service.FoodRecognitionService`
    - Implement `NewFoodRecognitionHandler(cfg config.DashScope) *FoodRecognitionHandler`
    - Implement `RecognizeFood(c *gin.Context)` handler: bind JSON request, validate image format and size, call service, return response envelope (`code`, `message`, `data`) matching existing API patterns
    - Handle all error cases from the design error handling table (400 for bad input, 502/504 for DashScope errors, 500 for config/parse errors, 200 with empty foods for no recognition)
    - _Requirements: 4.1, 4.2, 4.3, 4.4, 4.5, 4.6, 6.2, 6.3_

  - [x] 5.2 Register route in `internal/router/router.go`
    - Initialize `FoodRecognitionHandler` with DashScope config (update `SetupRouter` signature to accept `config.DashScope` or full `*config.Config`)
    - Add `foods.POST("/recognize", foodRecognitionHandler.RecognizeFood)` to the existing `foods` route group inside the `authorized` block
    - Update `cmd/main.go` to pass DashScope config to `SetupRouter`
    - _Requirements: 4.1, 4.2, 4.3_

  - [ ]* 5.3 Write unit tests for handler response envelope
    - Test success response format (code 200, message, data with foods array)
    - Test error responses: missing image (400), invalid format (400), oversized image (400)
    - Mock the service layer to test handler logic in isolation
    - _Requirements: 4.4, 4.5, 4.6_

- [x] 6. Implement frontend API client and recognition logic
  - [x] 6.1 Add `recognizeFood` API function to `src/api/diet.js`
    - Export `recognizeFood(imageBase64)` that calls `post('/foods/recognize', { image: imageBase64 })`
    - _Requirements: 4.1, 5.1_

  - [x] 6.2 Add AI recognition UI and logic to `src/pages/diet/add.vue`
    - Add data properties: `showRecognition` (boolean), `recognizing` (boolean), `recognizedFoods` (array), `recognitionError` (string)
    - Add "拍照识别" button with camera icon next to the "自定义食物" link in the food search card
    - Implement `startRecognition()` method: call `uni.chooseImage` (count: 1, sourceType: ['album', 'camera']), read file as base64, check size ≤ 2MB, call `recognizeFood` API
    - Handle camera permission denial: detect via error callback, show modal guiding to Settings
    - Show loading overlay with "AI 识别中..." text while `recognizing` is true
    - On success: populate `recognizedFoods` array, open recognition result bottom sheet (`showRecognition = true`)
    - On error: show toast with error message, set `recognitionError` for retry UI
    - _Requirements: 1.1, 1.2, 1.3, 1.4, 6.1, 6.2, 6.4_

  - [x] 6.3 Implement recognition result sheet in `src/pages/diet/add.vue`
    - Add bottom sheet (same pattern as portion-sheet and custom-sheet) showing recognized foods list
    - Each item shows: food name, confidence indicator (low confidence < 0.6 shows "结果仅供参考" warning), estimated grams (editable input), real-time nutrition preview (calories, protein, carbs, fat recalculated on gram change)
    - Add per-item checkbox for selective addition and per-item delete button
    - Add "全部添加" confirm button that converts selected recognized foods to `selectedFoods` format using `convertToSelectedFood()` and appends to `selectedFoods` array
    - Implement `convertToSelectedFood(recognizedFood)` — compute `calories`, `protein`, `carbs`, `fat` from per-100g values × grams / 100, rounded to 1 decimal
    - Close sheet after confirmation, update total nutrition summary
    - _Requirements: 3.1, 3.2, 3.3, 3.4, 3.5, 5.1, 5.2, 5.3, 6.3_

  - [x] 6.4 Add styles for AI recognition components in `src/pages/diet/add.vue`
    - Style the AI recognition button, loading overlay, recognition result sheet, confidence indicators, and editable gram inputs
    - Follow existing SCSS patterns and variables from the page
    - _Requirements: 1.1, 1.3, 3.1_

- [x] 7. Checkpoint — Ensure frontend builds and backend compiles
  - Ensure all tests pass, ask the user if questions arise.

- [ ] 8. Frontend property-based tests for nutrition calculations
  - [ ]* 8.1 Write property test: nutrition calculation correctness (Property 2)
    - **Property 2: Nutrition calculation correctness**
    - Use fast-check to generate arbitrary `RecognizedFood` objects with non-negative per-100g values and positive gram amounts
    - Assert `convertToSelectedFood` produces `calories` equal to `calories_per_100g * grams / 100` rounded to 1 decimal, same for protein, carbs, fat
    - Minimum 100 iterations
    - **Validates: Requirements 3.3, 5.1**

  - [ ]* 8.2 Write property test: total nutrition is sum of individual items (Property 3)
    - **Property 3: Total nutrition is sum of individual items**
    - Use fast-check to generate arbitrary arrays of selected food items (each with calories, protein, carbs, fat)
    - Assert computed `totalCalories` equals sum of individual `calories`, same for protein, carbs, fat
    - Minimum 100 iterations
    - **Validates: Requirements 3.4, 5.2**

  - [ ]* 8.3 Write property test: low confidence detection (Property 5)
    - **Property 5: Low confidence detection**
    - Use fast-check to generate `RecognizedFood` objects with confidence in [0, 1]
    - Assert food is flagged as low-confidence if and only if `confidence < 0.6`
    - Assert the "结果仅供参考" warning is triggered when at least one food in the list has low confidence
    - Minimum 100 iterations
    - **Validates: Requirements 6.3**

- [x] 9. Final checkpoint — Ensure all tests pass
  - Ensure all tests pass, ask the user if questions arise.

## Notes

- Tasks marked with `*` are optional and can be skipped for faster MVP
- Each task references specific requirements for traceability
- Checkpoints ensure incremental validation
- Property tests validate the 5 correctness properties defined in the design document (Properties 1 & 4 in Go backend, Properties 2, 3 & 5 in frontend with fast-check)
- Unit tests validate specific examples and edge cases
- The existing `selectedFoods` array format and `save()` flow are reused — AI-recognized foods are indistinguishable from manually-added foods after conversion
