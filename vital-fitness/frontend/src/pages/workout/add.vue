<template>
	<view class="container">
		<!-- 顶部导航栏 -->
		<u-navbar title="添加训练记录" fixed placeholder>
			<view slot="left">
				<u-icon name="arrow-left" size="20" @click="goBack"></u-icon>
			</view>
			<view slot="right">
				<text class="nav-btn" @click="saveWorkout">保存</text>
			</view>
		</u-navbar>

		<view class="form-container">
			<!-- 训练日期和时间 -->
			<view class="form-item">
				<text class="label">训练日期</text>
				<u-datetime-picker
					:show="showDatePicker"
					v-model="workoutForm.workout_date"
					mode="date"
					@confirm="confirmDate"
					@cancel="showDatePicker = false"
				></u-datetime-picker>
				<u-cell
					title="训练日期"
					:value="formatDate(workoutForm.workout_date)"
					isLink
					@click="showDatePicker = true"
				></u-cell>
			</view>

			<!-- 选择动作 -->
			<view class="form-item">
				<text class="label">训练动作</text>
				<u-cell
					title="选择动作"
					:value="selectedExercise?.name || '请选择'"
					isLink
					@click="showExercisePicker = true"
				></u-cell>

				<!-- 动作选择弹窗 -->
				<u-popup :show="showExercisePicker" mode="bottom" @close="showExercisePicker = false">
					<view class="popup-content">
						<view class="popup-header">
							<text class="popup-title">选择训练动作</text>
							<u-icon name="close" size="20" @click="showExercisePicker = false"></u-icon>
						</view>

						<!-- 动作分类筛选 -->
						<view class="category-filter">
							<u-scroll-list :indicator="false">
								<view class="category-item" :class="{ active: currentCategory === category }"
									v-for="category in categories" :key="category" @click="selectCategory(category)">
									{{ category }}
								</view>
							</u-scroll-list>
						</view>

						<!-- 动作列表 -->
						<view class="exercise-list">
							<view class="exercise-item" v-for="exercise in filteredExercises" :key="exercise.id"
								@click="selectExercise(exercise)">
								<text class="exercise-name">{{ exercise.name }}</text>
								<u-icon name="checkmark" v-if="selectedExercise?.id === exercise.id" color="#3c9cff"></u-icon>
							</view>
						</view>

						<!-- 添加自定义动作 -->
						<view class="add-custom-exercise">
							<u-button type="primary" icon="plus" @click="showCustomExercise = true">添加自定义动作</u-button>
						</view>
					</view>
				</u-popup>

				<!-- 添加自定义动作弹窗 -->
				<u-popup :show="showCustomExercise" mode="center" round closeable @close="showCustomExercise = false">
					<view class="custom-exercise-form">
						<view class="popup-header">
							<text class="popup-title">添加自定义动作</text>
						</view>

						<u-form :model="customExerciseForm" ref="customExerciseFormRef">
							<u-form-item label="动作名称" prop="name">
								<u-input v-model="customExerciseForm.name" placeholder="请输入动作名称" />
							</u-form-item>

							<u-form-item label="动作分类" prop="category">
								<u-picker :show="showCategoryPicker" :columns="categoryColumns" @confirm="confirmCategory"
									@cancel="showCategoryPicker = false"></u-picker>
								<u-cell title="选择分类" :value="customExerciseForm.category" isLink
									@click="showCategoryPicker = true"></u-cell>
							</u-form-item>

							<u-form-item label="动作描述" prop="description">
								<u-textarea v-model="customExerciseForm.description" placeholder="请输入动作描述（可选）" />
							</u-form-item>
						</u-form>

						<u-button type="primary" @click="saveCustomExercise">保存</u-button>
					</view>
				</u-popup>
			</view>

			<!-- 训练数据 -->
			<view class="form-item">
				<text class="label">训练数据</text>

				<view class="data-inputs">
					<view class="input-group">
						<text class="input-label">重量(kg)</text>
						<u-number-box v-model="workoutForm.weight" :min="0" :step="0.5"></u-number-box>
					</view>

					<view class="input-group">
						<text class="input-label">组数</text>
						<u-number-box v-model="workoutForm.sets" :min="1" :max="20"></u-number-box>
					</view>

					<view class="input-group">
						<text class="input-label">次数</text>
						<u-number-box v-model="workoutForm.reps" :min="1" :max="100"></u-number-box>
					</view>
				</view>
			</view>

			<!-- 休息时间和感受 -->
			<view class="form-item">
				<text class="label">其他信息</text>

				<view class="other-inputs">
					<view class="input-row">
						<text class="input-label">休息时间(秒)</text>
						<u-number-box v-model="workoutForm.rest_time" :min="0" :max="600"></u-number-box>
					</view>

					<view class="input-row">
						<text class="input-label">训练感受</text>
						<u-rate v-model="workoutForm.feeling" size="20" active-color="#ff9900"></u-rate>
					</view>
				</view>
			</view>

			<!-- 备注 -->
			<view class="form-item">
				<text class="label">备注</text>
				<u-textarea v-model="workoutForm.notes" placeholder="记录训练感受或其他备注信息..." />
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				showDatePicker: false,
				showExercisePicker: false,
				showCustomExercise: false,
				showCategoryPicker: false,

				workoutForm: {
					workout_date: Date.now(),
					exercise_id: null,
					weight: 0,
					sets: 1,
					reps: 1,
					rest_time: 60,
					feeling: 3,
					notes: ''
				},

				customExerciseForm: {
					name: '',
					category: '',
					description: ''
				},

				selectedExercise: null,

				categories: ['全部', '胸部', '背部', '腿部', '肩部', '手臂', '核心'],
				currentCategory: '全部',

				exercises: [
					{ id: 1, name: '杠铃卧推', category: '胸部' },
					{ id: 2, name: '哑铃飞鸟', category: '胸部' },
					{ id: 3, name: '俯卧撑', category: '胸部' },
					{ id: 4, name: '引体向上', category: '背部' },
					{ id: 5, name: '杠铃划船', category: '背部' },
					{ id: 6, name: '哑铃划船', category: '背部' },
					{ id: 7, name: '杠铃深蹲', category: '腿部' },
					{ id: 8, name: '腿举', category: '腿部' },
					{ id: 9, name: '腿弯举', category: '腿部' },
					{ id: 10, name: '杠铃推举', category: '肩部' },
					{ id: 11, name: '哑铃侧平举', category: '肩部' },
					{ id: 12, name: '哑铃弯举', category: '手臂' },
					{ id: 13, name: '杠铃弯举', category: '手臂' },
					{ id: 14, name: '平板支撑', category: '核心' }
				],

				categoryColumns: [['胸部', '背部', '腿部', '肩部', '手臂', '核心']]
			}
		},
		computed: {
			filteredExercises() {
				if (this.currentCategory === '全部') {
					return this.exercises
				}
				return this.exercises.filter(ex => ex.category === this.currentCategory)
			}
		},
		methods: {
			goBack() {
				uni.navigateBack()
			},

			saveWorkout() {
				if (!this.selectedExercise) {
					uni.showToast({
						title: '请选择训练动作',
						icon: 'none'
					})
					return
				}

				if (this.workoutForm.weight <= 0) {
					uni.showToast({
						title: '请填写重量',
						icon: 'none'
					})
					return
				}

				// 这里应该调用API保存训练记录
				uni.showToast({
					title: '保存成功',
					icon: 'success'
				})

				// 返回上一页
				setTimeout(() => {
					uni.navigateBack()
				}, 1000)
			},

			confirmDate(e) {
				this.workoutForm.workout_date = e.value
				this.showDatePicker = false
			},

			selectCategory(category) {
				this.currentCategory = category
			},

			selectExercise(exercise) {
				this.selectedExercise = exercise
				this.workoutForm.exercise_id = exercise.id
				this.showExercisePicker = false
			},

			confirmCategory(e) {
				this.customExerciseForm.category = e.value[0]
				this.showCategoryPicker = false
			},

			saveCustomExercise() {
				if (!this.customExerciseForm.name) {
					uni.showToast({
						title: '请输入动作名称',
						icon: 'none'
					})
					return
				}

				if (!this.customExerciseForm.category) {
					uni.showToast({
						title: '请选择动作分类',
						icon: 'none'
					})
					return
				}

				// 这里应该调用API保存自定义动作
				const newExercise = {
					id: this.exercises.length + 1,
					name: this.customExerciseForm.name,
					category: this.customExerciseForm.category
				}

				this.exercises.push(newExercise)
				this.selectExercise(newExercise)

				// 重置表单
				this.customExerciseForm = {
					name: '',
					category: '',
					description: ''
				}

				this.showCustomExercise = false

				uni.showToast({
					title: '添加成功',
					icon: 'success'
				})
			},

			formatDate(timestamp) {
				const date = new Date(timestamp)
				return `${date.getFullYear()}-${(date.getMonth() + 1).toString().padStart(2, '0')}-${date.getDate().toString().padStart(2, '0')}`
			}
		}
	}
</script>

<style lang="scss" scoped>
.container {
	padding: 20rpx;
}

.nav-btn {
	color: #3c9cff;
	font-size: 28rpx;
}

.form-container {
	.form-item {
		background: white;
		border-radius: 20rpx;
		padding: 0 30rpx 30rpx;
		margin-bottom: 20rpx;

		.label {
			display: block;
			font-size: 32rpx;
			font-weight: bold;
			color: #333;
			padding: 30rpx 0 20rpx;
		}

		.data-inputs {
			display: grid;
			grid-template-columns: repeat(3, 1fr);
			gap: 20rpx;

			.input-group {
				text-align: center;

				.input-label {
					display: block;
					font-size: 24rpx;
					color: #666;
					margin-bottom: 10rpx;
				}
			}
		}

		.other-inputs {
			.input-row {
				display: flex;
				align-items: center;
				justify-content: space-between;
				margin-bottom: 30rpx;

				&:last-child {
					margin-bottom: 0;
				}

				.input-label {
					font-size: 28rpx;
					color: #333;
				}
			}
		}
	}
}

.popup-content {
	.popup-header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: 30rpx;
		border-bottom: 1rpx solid #eee;

		.popup-title {
			font-size: 32rpx;
			font-weight: bold;
		}
	}

	.category-filter {
		padding: 20rpx 0;
		border-bottom: 1rpx solid #eee;

		.category-item {
			padding: 10rpx 20rpx;
			margin: 0 10rpx;
			border-radius: 10rpx;
			background: #f5f5f5;
			white-space: nowrap;

			&.active {
				background: #3c9cff;
				color: white;
			}
		}
	}

	.exercise-list {
		max-height: 500rpx;
		overflow-y: auto;

		.exercise-item {
			display: flex;
			justify-content: space-between;
			align-items: center;
			padding: 30rpx;
			border-bottom: 1rpx solid #eee;

			.exercise-name {
				font-size: 28rpx;
				color: #333;
			}
		}
	}

	.add-custom-exercise {
		padding: 30rpx;
		text-align: center;
		border-top: 1rpx solid #eee;
	}
}

.custom-exercise-form {
	padding: 30rpx;
	width: 600rpx;

	.popup-header {
		text-align: center;
		margin-bottom: 30rpx;

		.popup-title {
			font-size: 32rpx;
			font-weight: bold;
		}
	}

	.u-form-item {
		margin-bottom: 30rpx;
	}
}
</style>