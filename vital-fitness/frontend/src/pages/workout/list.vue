<template>
	<view class="container">
		<view class="nav-bar">
			<text class="nav-title">训练记录</text>
			<text class="add-btn" @tap="addWorkout">+ 记录</text>
		</view>

		<!-- 分类筛选 -->
		<view class="filter-bar">
			<text
				v-for="(tab, index) in tabList"
				:key="index"
				class="filter-tab"
				:class="{ active: currentTab === index }"
				@tap="changeTab(index)"
			>{{ tab.name }}</text>
		</view>

		<!-- 训练记录列表 -->
		<view class="workout-list">
			<view v-for="(workout, index) in workoutList" :key="index" class="workout-item">
				<view class="workout-header">
					<text class="workout-date">{{ formatDate(workout.workout_date) }}</text>
					<text class="workout-time">{{ formatTime(workout.created_at) }}</text>
				</view>
				<view class="workout-content">
					<text class="exercise-name">{{ workout.exercise.name }}</text>
					<view class="workout-details">
						<text class="detail-tag">{{ workout.weight }}kg</text>
						<text class="detail-tag" v-if="workout.sets > 1">{{ workout.sets }}组</text>
						<text class="detail-tag" v-if="workout.reps > 1">{{ workout.reps }}次</text>
					</view>
				</view>
				<text v-if="workout.notes" class="workout-notes">{{ workout.notes }}</text>
			</view>

			<!-- 空状态 -->
			<view v-if="workoutList.length === 0" class="empty-state">
				<text class="empty-text">暂无训练记录</text>
				<button class="add-first-btn" type="primary" @tap="addWorkout">添加第一条记录</button>
			</view>
		</view>
	</view>
</template>

<script>
	import { getWorkouts } from '../../api/workout'

	export default {
		data() {
			return {
				tabList: [
					{ name: '全部' }, { name: '胸部' }, { name: '背部' },
					{ name: '腿部' }, { name: '肩部' }, { name: '手臂' }
				],
				currentTab: 0,
				workoutList: []
			}
		},
		onShow() {
			this.loadWorkouts()
		},
		methods: {
			changeTab(index) {
				this.currentTab = index
				this.loadWorkouts()
			},
			async loadWorkouts() {
				try {
					const category = this.tabList[this.currentTab].name
					const res = await getWorkouts({
						category: category === '全部' ? '' : category,
						page: 1, page_size: 50
					})
					this.workoutList = res.data || []
				} catch (e) {}
			},
			addWorkout() {
				uni.navigateTo({ url: '/pages/workout/add' })
			},
			formatDate(dateStr) {
				if (!dateStr) return ''
				const date = new Date(dateStr)
				const today = new Date()
				if (date.toDateString() === today.toDateString()) return '今天'
				const yesterday = new Date(today)
				yesterday.setDate(yesterday.getDate() - 1)
				if (date.toDateString() === yesterday.toDateString()) return '昨天'
				return (date.getMonth() + 1) + '月' + date.getDate() + '日'
			},
			formatTime(dateStr) {
				if (!dateStr) return ''
				const d = new Date(dateStr)
				return String(d.getHours()).padStart(2, '0') + ':' + String(d.getMinutes()).padStart(2, '0')
			}
		}
	}
</script>

<style lang="scss" scoped>
.container { padding: 20rpx; }

.nav-bar { display: flex; justify-content: space-between; align-items: center; margin-bottom: 15rpx;
	.nav-title { font-size: 36rpx; font-weight: 700; color: #1c1c1e; }
}

.add-btn {
	color: #3c9cff;
	font-size: 28rpx;
	font-weight: bold;
}

.filter-bar {
	display: flex;
	flex-wrap: wrap;
	gap: 15rpx;
	background: white;
	border-radius: 20rpx;
	padding: 20rpx;
	margin-bottom: 20rpx;

	.filter-tab {
		padding: 10rpx 24rpx;
		border-radius: 30rpx;
		background: #f5f5f5;
		font-size: 24rpx;
		color: #666;

		&.active { background: #3c9cff; color: white; }
	}
}

.workout-list {
	.workout-item {
		background: white;
		border-radius: 20rpx;
		padding: 30rpx;
		margin-bottom: 15rpx;

		.workout-header {
			display: flex;
			justify-content: space-between;
			margin-bottom: 15rpx;

			.workout-date { font-size: 24rpx; color: #999; }
			.workout-time { font-size: 24rpx; color: #999; }
		}

		.workout-content {
			.exercise-name {
				font-size: 32rpx;
				font-weight: bold;
				color: #333;
				display: block;
				margin-bottom: 15rpx;
			}

			.workout-details {
				display: flex;
				gap: 15rpx;

				.detail-tag {
					padding: 8rpx 20rpx;
					background: #f0f7ff;
					color: #3c9cff;
					border-radius: 10rpx;
					font-size: 24rpx;
				}
			}
		}

		.workout-notes {
			margin-top: 15rpx;
			font-size: 24rpx;
			color: #999;
		}
	}

	.empty-state {
		text-align: center;
		padding: 100rpx 0;
		background: white;
		border-radius: 20rpx;

		.empty-text { font-size: 28rpx; color: #999; display: block; margin-bottom: 30rpx; }
		.add-first-btn { width: 60%; font-size: 28rpx; }
	}
}
</style>
