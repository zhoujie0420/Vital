<template>
	<view class="container">
		<!-- 顶部导航栏 -->
		<u-navbar title="训练记录" fixed placeholder>
			<view slot="right">
				<u-icon name="plus" size="20" @click="addWorkout"></u-icon>
			</view>
		</u-navbar>

		<!-- 筛选条件 -->
		<view class="filter-bar">
			<u-tabs :list="tabList" :current="currentTab" @change="changeTab"></u-tabs>
		</view>

		<!-- 训练记录列表 -->
		<view class="workout-list">
			<view v-for="(workout, index) in workoutList" :key="index" class="workout-item" @click="viewWorkout(workout.id)">
				<view class="workout-header">
					<text class="workout-date">{{ formatDate(workout.workout_date) }}</text>
					<text class="workout-time">{{ formatTime(workout.created_at) }}</text>
				</view>

				<view class="workout-content">
					<view class="exercise-info">
						<u-icon name="level" size="18" color="#3c9cff"></u-icon>
						<text class="exercise-name">{{ workout.exercise.name }}</text>
					</view>

					<view class="workout-details">
						<text class="detail-item">重量: {{ workout.weight }}kg</text>
						<text class="detail-item">组数: {{ workout.sets }}组</text>
						<text class="detail-item">次数: {{ workout.reps }}次</text>
					</view>

					<view class="workout-feeling">
						<text class="feeling-label">感受:</text>
						<u-rate :value="workout.feeling" readonly size="14" active-color="#ff9900"></u-rate>
					</view>
				</view>

				<view class="workout-footer">
					<u-tag :text="workout.exercise.category" type="primary" size="mini" />
					<text v-if="workout.notes" class="notes-preview">{{ workout.notes }}</text>
				</view>
			</view>

			<!-- 空状态 -->
			<view v-if="workoutList.length === 0" class="empty-state">
				<u-empty text="暂无训练记录" mode="list"></u-empty>
				<u-button type="primary" icon="plus" @click="addWorkout" style="margin-top: 30rpx;">添加第一条记录</u-button>
			</view>
		</view>

		<!-- 加载更多 -->
		<u-loadmore :status="loadStatus" v-if="workoutList.length > 0" />
	</view>
</template>

<script>
	import { getWorkouts } from '../../api/workout'

	export default {
		data() {
			return {
				tabList: [
					{ name: '全部' },
					{ name: '胸部' },
					{ name: '背部' },
					{ name: '腿部' },
					{ name: '肩部' }
				],
				currentTab: 0,
				workoutList: [],
				loadStatus: 'loadmore',
				page: 1
			}
		},
		onShow() {
			this.page = 1
			this.loadWorkouts()
		},
		methods: {
			changeTab(index) {
				this.currentTab = index
				this.page = 1
				this.loadWorkouts()
			},

			async loadWorkouts() {
				this.loadStatus = 'loading'
				try {
					const category = this.tabList[this.currentTab].name
					const res = await getWorkouts({
						category: category === '全部' ? '' : category,
						page: this.page,
						page_size: 20
					})
					this.workoutList = res.data || []
					this.loadStatus = this.workoutList.length < 20 ? 'nomore' : 'loadmore'
				} catch (e) {
					this.loadStatus = 'loadmore'
				}
			},

			addWorkout() {
				uni.navigateTo({
					url: '/pages/workout/add'
				})
			},

			viewWorkout(id) {
				uni.navigateTo({
					url: `/pages/workout/detail?id=${id}`
				})
			},

			formatDate(dateString) {
				const date = new Date(dateString)
				const today = new Date()

				if (date.toDateString() === today.toDateString()) {
					return '今天'
				}

				const yesterday = new Date(today)
				yesterday.setDate(yesterday.getDate() - 1)

				if (date.toDateString() === yesterday.toDateString()) {
					return '昨天'
				}

				return `${date.getMonth() + 1}月${date.getDate()}日`
			},

			formatTime(dateTimeString) {
				const date = new Date(dateTimeString)
				return `${date.getHours().toString().padStart(2, '0')}:${date.getMinutes().toString().padStart(2, '0')}`
			}
		}
	}
</script>

<style lang="scss" scoped>
.container {
	padding: 20rpx;
}

.filter-bar {
	background: white;
	border-radius: 20rpx;
	margin-bottom: 20rpx;
	overflow: hidden;
}

.workout-list {
	.workout-item {
		background: white;
		border-radius: 20rpx;
		padding: 30rpx;
		margin-bottom: 20rpx;

		.workout-header {
			display: flex;
			justify-content: space-between;
			align-items: center;
			margin-bottom: 20rpx;

			.workout-date {
				font-size: 32rpx;
				font-weight: bold;
				color: #333;
			}

			.workout-time {
				font-size: 24rpx;
				color: #999;
			}
		}

		.workout-content {
			.exercise-info {
				display: flex;
				align-items: center;
				margin-bottom: 20rpx;

				.exercise-name {
					font-size: 28rpx;
					font-weight: bold;
					color: #333;
					margin-left: 10rpx;
				}
			}

			.workout-details {
				display: flex;
				flex-wrap: wrap;
				gap: 20rpx;
				margin-bottom: 20rpx;

				.detail-item {
					font-size: 26rpx;
					color: #666;
					background: #f5f5f5;
					padding: 10rpx 20rpx;
					border-radius: 10rpx;
				}
			}

			.workout-feeling {
				display: flex;
				align-items: center;

				.feeling-label {
					font-size: 26rpx;
					color: #666;
					margin-right: 10rpx;
				}
			}
		}

		.workout-footer {
			margin-top: 20rpx;
			padding-top: 20rpx;
			border-top: 1rpx solid #eee;

			.notes-preview {
				font-size: 24rpx;
				color: #999;
				margin-left: 10rpx;
			}
		}
	}

	.empty-state {
		text-align: center;
		padding: 100rpx 0;
		background: white;
		border-radius: 20rpx;
	}
}
</style>