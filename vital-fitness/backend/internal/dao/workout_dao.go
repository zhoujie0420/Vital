package dao

import (
	"vital-fitness/backend/internal/model"
	"vital-fitness/backend/internal/utils"

	"gorm.io/gorm"
)

type WorkoutDAO struct {
	db *gorm.DB
}

func NewWorkoutDAO() *WorkoutDAO {
	return &WorkoutDAO{db: utils.GetDB()}
}

func (d *WorkoutDAO) Create(workout *model.Workout) error {
	return d.db.Create(workout).Error
}

func (d *WorkoutDAO) FindByID(id uint, userID uint) (*model.Workout, error) {
	var workout model.Workout
	err := d.db.Preload("Exercise").Where("id = ? AND user_id = ?", id, userID).First(&workout).Error
	return &workout, err
}

func (d *WorkoutDAO) FindByUserID(userID uint, category string, page, pageSize int) ([]model.Workout, int64, error) {
	var workouts []model.Workout
	var total int64

	query := d.db.Where("user_id = ?", userID)
	if category != "" && category != "全部" {
		query = query.Joins("JOIN exercises ON exercises.id = workouts.exercise_id").
			Where("exercises.category = ?", category)
	}

	query.Model(&model.Workout{}).Count(&total)
	err := query.Preload("Exercise").
		Order("workout_date DESC, created_at DESC").
		Offset((page - 1) * pageSize).Limit(pageSize).
		Find(&workouts).Error

	return workouts, total, err
}

func (d *WorkoutDAO) Update(workout *model.Workout) error {
	return d.db.Save(workout).Error
}

func (d *WorkoutDAO) Delete(id uint, userID uint) error {
	return d.db.Where("id = ? AND user_id = ?", id, userID).Delete(&model.Workout{}).Error
}

// Exercise 相关
func (d *WorkoutDAO) FindAllExercises(category string) ([]model.Exercise, error) {
	var exercises []model.Exercise
	query := d.db.Model(&model.Exercise{})
	if category != "" && category != "全部" {
		query = query.Where("category = ?", category)
	}
	err := query.Order("category, name").Find(&exercises).Error
	return exercises, err
}

func (d *WorkoutDAO) CreateExercise(exercise *model.Exercise) error {
	return d.db.Create(exercise).Error
}
