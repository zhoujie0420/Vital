# Requirements Document

## Introduction

AI 食物拍照识别功能，允许用户在饮食记录页面通过拍照或从相册选择食物图片，由后端调用 AI 视觉模型自动识别食物种类并估算营养信息（热量、蛋白质、碳水化合物、脂肪），识别结果可直接添加到饮食记录中。该功能集成到现有的饮食记录添加流程（diet/add.vue），作为手动搜索食物的补充入口。

## Glossary

- **Recognition_Service**: 后端 AI 食物识别服务，负责接收图片、调用阿里通义千问视觉模型（qwen-vl）API、解析返回结果
- **Food_Camera_UI**: 前端拍照/选图交互组件，集成在饮食记录添加页面
- **Recognition_Result**: AI 模型返回的识别结果数据结构，包含食物名称、估算份量、营养信息
- **Image_Uploader**: 前端图片上传模块，负责图片压缩和上传至后端
- **Nutrition_Estimator**: AI 模型中负责根据食物种类和估算份量计算营养数据的逻辑

## Requirements

### Requirement 1: 食物图片拍摄与选择

**User Story:** As a 用户, I want 在记录饮食时通过拍照或从相册选择食物图片, so that 我可以快速输入食物信息而无需手动搜索。

#### Acceptance Criteria

1. WHEN 用户点击拍照识别按钮, THE Food_Camera_UI SHALL 调用微信小程序相机 API 打开拍照界面或相册选择界面
2. WHEN 用户完成拍照或选择图片, THE Image_Uploader SHALL 将图片压缩至不超过 2MB 后上传至后端
3. WHILE 图片正在上传和识别中, THE Food_Camera_UI SHALL 显示加载动画和"识别中"提示文字
4. IF 用户未授权相机权限, THEN THE Food_Camera_UI SHALL 显示权限引导提示并提供跳转设置页的入口

### Requirement 2: AI 食物识别与营养估算

**User Story:** As a 用户, I want 上传的食物图片被自动识别出食物种类和营养信息, so that 我不需要手动查找和计算每种食物的热量。

#### Acceptance Criteria

1. WHEN Recognition_Service 接收到食物图片, THE Recognition_Service SHALL 调用阿里通义千问视觉模型（qwen-vl）API 进行食物识别
2. WHEN AI 模型返回识别结果, THE Recognition_Service SHALL 解析结果并返回包含食物名称、估算克数、每 100g 热量、蛋白质、碳水化合物、脂肪的结构化数据
3. THE Recognition_Service SHALL 在接收图片后 15 秒内返回识别结果
4. WHEN 图片中包含多种食物, THE Recognition_Service SHALL 分别识别每种食物并返回各自的营养信息列表
5. IF AI 模型无法识别图片中的食物, THEN THE Recognition_Service SHALL 返回明确的错误提示信息，说明无法识别

### Requirement 3: 识别结果展示与确认

**User Story:** As a 用户, I want 查看 AI 识别出的食物列表和营养信息并进行确认或修改, so that 我可以确保记录的数据准确。

#### Acceptance Criteria

1. WHEN 识别结果返回成功, THE Food_Camera_UI SHALL 以列表形式展示每种识别出的食物名称、估算克数和对应热量
2. WHEN 用户查看识别结果, THE Food_Camera_UI SHALL 为每项食物提供修改份量的输入框
3. WHEN 用户修改某项食物的份量, THE Food_Camera_UI SHALL 实时重新计算该食物的热量、蛋白质、碳水化合物和脂肪数值
4. WHEN 用户确认识别结果, THE Food_Camera_UI SHALL 将选中的食物项添加到当前饮食记录的已选食物列表中
5. WHEN 用户对某项识别结果不满意, THE Food_Camera_UI SHALL 提供删除该项食物的操作

### Requirement 4: 后端 API 接口

**User Story:** As a 开发者, I want 后端提供标准化的食物识别 API 接口, so that 前端可以通过统一的方式调用 AI 识别功能。

#### Acceptance Criteria

1. THE Recognition_Service SHALL 提供 POST /api/v1/foods/recognize 接口接收图片数据
2. THE Recognition_Service SHALL 要求请求携带有效的 JWT 认证令牌
3. WHEN 请求缺少有效认证令牌, THE Recognition_Service SHALL 返回 401 状态码
4. WHEN 上传的文件不是有效的图片格式（JPG、PNG）, THE Recognition_Service SHALL 返回 400 状态码和格式错误提示
5. WHEN 上传的图片大小超过 5MB, THE Recognition_Service SHALL 返回 400 状态码和文件过大提示
6. THE Recognition_Service SHALL 返回 JSON 格式的响应，包含 code、message 和 data 字段，与现有 API 响应格式一致

### Requirement 5: 识别结果与现有饮食记录系统集成

**User Story:** As a 用户, I want AI 识别的食物可以直接保存到我的饮食记录中, so that 拍照识别的结果能无缝融入现有的饮食追踪流程。

#### Acceptance Criteria

1. WHEN 用户确认 AI 识别的食物列表, THE Food_Camera_UI SHALL 将食物数据转换为与手动添加食物相同的数据格式（name、grams、calories、protein、carbs、fat）
2. WHEN AI 识别的食物被添加到已选列表, THE Food_Camera_UI SHALL 更新页面底部的总热量和宏量营养素汇总数据
3. THE Food_Camera_UI SHALL 允许用户在同一次饮食记录中混合使用 AI 识别和手动搜索两种方式添加食物

### Requirement 6: 错误处理与降级

**User Story:** As a 用户, I want 在 AI 识别失败时获得清晰的反馈和替代方案, so that 我仍然可以完成饮食记录。

#### Acceptance Criteria

1. IF 网络请求超时（超过 15 秒）, THEN THE Food_Camera_UI SHALL 显示超时提示并提供重试按钮
2. IF 后端服务返回错误, THEN THE Food_Camera_UI SHALL 显示错误信息并引导用户使用手动搜索方式
3. IF AI 模型返回低置信度结果, THEN THE Recognition_Service SHALL 在响应中标注置信度信息，THE Food_Camera_UI SHALL 向用户显示"结果仅供参考"的提示
4. IF 图片模糊或光线不足导致识别失败, THEN THE Food_Camera_UI SHALL 提示用户重新拍摄并给出拍照建议（光线充足、食物居中）
