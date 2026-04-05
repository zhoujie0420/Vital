package service

import (
	"errors"

	"vital-fitness/backend/internal/dao"
	"vital-fitness/backend/internal/model"
)

type WorkoutService struct {
	workoutDAO *dao.WorkoutDAO
}

func NewWorkoutService() *WorkoutService {
	return &WorkoutService{workoutDAO: dao.NewWorkoutDAO()}
}

func (s *WorkoutService) CreateWorkout(userID uint, req *model.CreateWorkoutRequest) (*model.Workout, error) {
	workout := &model.Workout{
		UserID:      userID,
		ExerciseID:  req.ExerciseID,
		WorkoutDate: req.WorkoutDate,
		Weight:      req.Weight,
		Sets:        req.Sets,
		Reps:        req.Reps,
		RestTime:    req.RestTime,
		Feeling:     req.Feeling,
		Notes:       req.Notes,
	}
	if err := s.workoutDAO.Create(workout); err != nil {
		return nil, errors.New("创建训练记录失败")
	}
	return s.workoutDAO.FindByID(workout.ID, userID)
}

func (s *WorkoutService) GetWorkouts(userID uint, category string, page, pageSize int) ([]model.Workout, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}
	return s.workoutDAO.FindByUserID(userID, category, page, pageSize)
}

func (s *WorkoutService) GetWorkout(id uint, userID uint) (*model.Workout, error) {
	return s.workoutDAO.FindByID(id, userID)
}

func (s *WorkoutService) UpdateWorkout(id uint, userID uint, req *model.UpdateWorkoutRequest) (*model.Workout, error) {
	workout, err := s.workoutDAO.FindByID(id, userID)
	if err != nil {
		return nil, errors.New("训练记录不存在")
	}

	if req.ExerciseID != 0 {
		workout.ExerciseID = req.ExerciseID
	}
	if !req.WorkoutDate.IsZero() {
		workout.WorkoutDate = req.WorkoutDate
	}
	if req.Weight != 0 {
		workout.Weight = req.Weight
	}
	if req.Sets != 0 {
		workout.Sets = req.Sets
	}
	if req.Reps != 0 {
		workout.Reps = req.Reps
	}
	if req.RestTime != 0 {
		workout.RestTime = req.RestTime
	}
	if req.Feeling != 0 {
		workout.Feeling = req.Feeling
	}
	if req.Notes != "" {
		workout.Notes = req.Notes
	}

	if err := s.workoutDAO.Update(workout); err != nil {
		return nil, errors.New("更新失败")
	}
	return workout, nil
}

func (s *WorkoutService) DeleteWorkout(id uint, userID uint) error {
	return s.workoutDAO.Delete(id, userID)
}

func (s *WorkoutService) GetExercises(category string) ([]model.Exercise, error) {
	return s.workoutDAO.FindAllExercises(category)
}

func (s *WorkoutService) CreateExercise(req *model.Exercise) error {
	return s.workoutDAO.CreateExercise(req)
}
