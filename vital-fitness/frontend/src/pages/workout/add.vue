<template>
	<view class="container">
		<u-navbar title="记录训练" fixed placeholder>
			<view slot="left">
				<u-icon name="arrow-left" size="20" @click="goBack"></u-icon>
			</view>
		</u-navbar>

		<view class="form-container">
			<!-- 选择动作 -->
			<view class="form-item">
				<text class="label">选择动作</text>

				<!-- 分类筛选 -->
				<view class="category-tabs">
					<text
						v-for="cat in categories"
						:key="cat"
						class="cat-tab"
						:class="{ active: currentCategory === cat }"
						@tap="currentCategory = cat"
					>{{ cat }}</text>
				</view>

				<!-- 动作列表 -->
				<view class="exercise-grid">
					<view
						v-for="ex in filteredExercises"
						:key="ex.id"
						class="exercise-card"
						:class="{ selected: selectedExercise && selectedExercise.id === ex.id }"
						@tap="selectExercise(ex)"
					>
						<text class="exercise-name">{{ ex.name }}</text>
					</view>
				</view>
			</view>

			<!-- 重量输入 -->
			<view class="form-item" v-if="selectedExercise">
				<text class="label">重量 (kg)</text>
				<view class="weight-input">
					<text class="weight-btn" @tap="adjustWeight(-2.5)">-</text>
					<input
						type="digit"
						v-model="weight"
						class="weight-value"
						placeholder="0"
					/>
					<text class="weight-btn" @tap="adjustWeight(2.5)">+</text>
				</view>
				<view class="quick-weights">
					<text class="quick-btn" v-for="w in quickWeights" :key="w" @tap="weight = w">{{ w }}kg</text>
				</view>
			</view>

			<!-- 保存按钮 -->
			<view class="form-item" v-if="selectedExercise">
				<button class="save-btn" type="primary" :loading="saving" @tap="saveWorkout">
					{{ saving ? '保存中...' : '保存记录' }}
				</button>
			</view>
		</view>
	</view>
</template>

<script>
	import { getExercises, createWorkout } from '../../api/workout'

	export default {
		data() {
			return {
				saving: false,
				categories: ['全部', '胸部', '背部', '腿部', '肩部', '手臂', '核心'],
				currentCategory: '全部',
				exercises: [],
				selectedExercise: null,
				weight: '',
				quickWeights: [20, 40, 60, 80, 100]
			}
		},
		computed: {
			filteredExercises() {
				if (this.currentCategory === '全部') return this.exercises
				return this.exercises.filter(ex => ex.category === this.currentCategory)
			}
		},
		onLoad() {
			this.loadExercises()
		},
		methods: {
			async loadExercises() {
				try {
					const res = await getExercises('')
					this.exercises = res.data || []
				} catch (e) {}
			},
			goBack() { uni.navigateBack() },
			selectExercise(ex) { this.selectedExercise = ex },
			adjustWeight(delta) {
				const current = parseFloat(this.weight) || 0
				const next = current + delta
				if (next >= 0) this.weight = next
			},
			async saveWorkout() {
				if (!this.weight || parseFloat(this.weight) <= 0) {
					uni.showToast({ title: '请输入重量', icon: 'none' })
					return
				}
				this.saving = true
				try {
					await createWorkout({
						exercise_id: this.selectedExercise.id,
						workout_date: new Date().toISOString(),
						weight: parseFloat(this.weight),
						sets: 1,
						reps: 1
					})
					uni.showToast({ title: '保存成功', icon: 'success' })
					setTimeout(() => uni.navigateBack(), 1000)
				} catch (e) {
				} finally {
					this.saving = false
				}
			}
		}
	}
</script>

<style lang="scss" scoped>
.container {
	padding: 20rpx;
}

.form-container {
	.form-item {
		background: white;
		border-radius: 20rpx;
		padding: 30rpx;
		margin-bottom: 20rpx;

		.label {
			font-size: 32rpx;
			font-weight: bold;
			color: #333;
			margin-bottom: 20rpx;
			display: block;
		}

		.category-tabs {
			display: flex;
			flex-wrap: wrap;
			gap: 15rpx;
			margin-bottom: 20rpx;

			.cat-tab {
				padding: 10rpx 24rpx;
				border-radius: 30rpx;
				background: #f5f5f5;
				font-size: 24rpx;
				color: #666;

				&.active {
					background: #3c9cff;
					color: white;
				}
			}
		}

		.exercise-grid {
			display: flex;
			flex-wrap: wrap;
			gap: 15rpx;

			.exercise-card {
				padding: 20rpx 30rpx;
				border-radius: 15rpx;
				background: #f5f5f5;
				border: 2rpx solid transparent;

				&.selected {
					background: #e8f4ff;
					border-color: #3c9cff;
				}

				.exercise-name {
					font-size: 26rpx;
					color: #333;
				}
			}
		}

		.weight-input {
			display: flex;
			align-items: center;
			justify-content: center;
			margin: 30rpx 0;

			.weight-btn {
				width: 80rpx;
				height: 80rpx;
				line-height: 80rpx;
				text-align: center;
				background: #3c9cff;
				color: white;
				border-radius: 50%;
				font-size: 40rpx;
				font-weight: bold;
			}

			.weight-value {
				width: 200rpx;
				text-align: center;
				font-size: 60rpx;
				font-weight: bold;
				color: #333;
				margin: 0 30rpx;
			}
		}

		.quick-weights {
			display: flex;
			justify-content: center;
			gap: 20rpx;

			.quick-btn {
				padding: 10rpx 24rpx;
				border-radius: 10rpx;
				background: #f5f5f5;
				font-size: 24rpx;
				color: #666;
			}
		}

		.save-btn {
			border-radius: 15rpx;
			font-size: 32rpx;
		}
	}
}
</style>
