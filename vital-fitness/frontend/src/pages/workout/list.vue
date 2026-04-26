<template>
	<view class="page" :style="{ paddingTop: topPadding + 'px' }">
		<view class="page-header">
			<text class="page-title">训练记录</text>
			<text class="header-action" @tap="addWorkout">+ 记录</text>
		</view>

		<!-- 分类筛选 -->
		<scroll-view scroll-x class="filter-scroll">
			<view class="filter-bar">
				<text v-for="(tab, index) in tabList" :key="index" class="filter-pill"
					:class="{ active: currentTab === index }" @tap="changeTab(index)">
					{{ tab.name }}
				</text>
			</view>
		</scroll-view>

		<!-- 按天分组的卡片 -->
		<view class="day-list">
			<view v-for="(day, i) in groupedDays" :key="i" class="day-card" @tap="goToDetail(day.date)">
				<view class="day-header">
					<view class="day-date-wrap">
						<text class="day-label">{{ day.label }}</text>
						<text class="day-date">{{ day.dateStr }}</text>
					</view>
					<view class="day-summary">
						<text class="day-count">{{ day.workouts.length }}个动作</text>
						<text class="day-arrow">›</text>
					</view>
				</view>
				<view class="day-exercises">
					<text v-for="(w, j) in day.workouts.slice(0, 4)" :key="j" class="exercise-tag">
						{{ w.exercise.name }}
					</text>
					<text v-if="day.workouts.length > 4" class="exercise-more">+{{ day.workouts.length - 4 }}</text>
				</view>
			</view>
		</view>

		<!-- 空状态 -->
		<view v-if="groupedDays.length === 0" class="empty">
			<text class="empty-icon">💪</text>
			<text class="empty-title">暂无训练记录</text>
			<text class="empty-desc">记录每一次训练，见证你的进步</text>
			<view class="empty-btn" @tap="addWorkout">
				<text>开始记录</text>
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
		onShow() { this.loadWorkouts() },
		computed: {
			topPadding() {
				const app = getApp()
				return (app.globalData?.customBarHeight || 88) + 8
			},
			groupedDays() {
				const map = {}
				this.workoutList.forEach(w => {
					const d = new Date(w.workout_date)
					const key = d.toISOString().split('T')[0]
					if (!map[key]) map[key] = { date: key, workouts: [] }
					map[key].workouts.push(w)
				})
				return Object.values(map).sort((a, b) => b.date.localeCompare(a.date)).map(day => {
					const d = new Date(day.date + 'T00:00:00')
					const today = new Date()
					const yesterday = new Date(today)
					yesterday.setDate(yesterday.getDate() - 1)
					let label = ''
					if (d.toDateString() === today.toDateString()) label = '今天'
					else if (d.toDateString() === yesterday.toDateString()) label = '昨天'
					else label = (d.getMonth() + 1) + '月' + d.getDate() + '日'
					return {
						...day,
						label,
						dateStr: d.getFullYear() + '.' + (d.getMonth()+1) + '.' + d.getDate()
					}
				})
			}
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
						page: 1, page_size: 200
					})
					this.workoutList = res.data || []
				} catch (e) {}
			},
			addWorkout() { uni.navigateTo({ url: '/pages/workout/add' }) },
			goToDetail(date) {
				uni.navigateTo({ url: '/pages/workout/detail?date=' + date })
			}
		}
	}
</script>

<style lang="scss" scoped>
.page {
	padding: 0 32rpx;
	padding-bottom: 40rpx;
	min-height: 100vh;
	background: linear-gradient(180deg, #eef4ff 0%, #f2f2f7 280rpx);
}

.page-header {
	display: flex;
	justify-content: space-between;
	align-items: baseline;
	margin-bottom: 24rpx;

	.page-title { font-size: 52rpx; font-weight: 700; color: #1c1c1e; letter-spacing: -1rpx; }
	.header-action { font-size: 30rpx; color: #007aff; font-weight: 600; }
}

.filter-scroll {
	white-space: nowrap;
	margin-bottom: 20rpx;

	.filter-bar {
		display: inline-flex;
		gap: 12rpx;

		.filter-pill {
			display: inline-block;
			padding: 12rpx 28rpx;
			border-radius: 24rpx;
			background: rgba(0, 0, 0, 0.04);
			font-size: 26rpx;
			color: #636366;
			font-weight: 500;

			&.active {
				background: #007aff;
				color: #fff;
				box-shadow: 0 4rpx 12rpx rgba(0, 122, 255, 0.3);
			}
			&:active { transform: scale(0.95); }
		}
	}
}

.day-card {
	background: #fff;
	border-radius: 20rpx;
	padding: 28rpx 28rpx 28rpx 22rpx;
	margin-bottom: 16rpx;
	box-shadow: 0 4rpx 20rpx rgba(0, 0, 0, 0.06);
	transition: transform 0.15s ease;
	border-left: 6rpx solid #0a84ff;

	&:active { transform: scale(0.98); }

	.day-header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: 16rpx;

		.day-date-wrap {
			.day-label {
				display: block;
				font-size: 32rpx;
				font-weight: 600;
				color: #1c1c1e;
			}
			.day-date {
				display: block;
				font-size: 22rpx;
				color: #8e8e93;
				margin-top: 4rpx;
			}
		}

		.day-summary {
			display: flex;
			align-items: center;
			gap: 8rpx;

			.day-count {
				font-size: 26rpx;
				color: #8e8e93;
				font-weight: 500;
			}
			.day-arrow {
				font-size: 36rpx;
				color: #c7c7cc;
				font-weight: 300;
			}
		}
	}

	.day-exercises {
		display: flex;
		flex-wrap: wrap;
		gap: 10rpx;

		.exercise-tag {
			padding: 8rpx 20rpx;
			background: rgba(0, 122, 255, 0.08);
			color: #007aff;
			border-radius: 12rpx;
			font-size: 24rpx;
			font-weight: 500;
		}
		.exercise-more {
			padding: 8rpx 16rpx;
			color: #8e8e93;
			font-size: 24rpx;
			font-weight: 500;
		}
	}
}

.empty {
	text-align: center;
	padding: 120rpx 0;

	.empty-icon { display: block; font-size: 96rpx; margin-bottom: 24rpx; }
	.empty-title { display: block; font-size: 34rpx; font-weight: 600; color: #1c1c1e; margin-bottom: 8rpx; }
	.empty-desc { display: block; font-size: 26rpx; color: #8e8e93; margin-bottom: 40rpx; }
	.empty-btn {
		display: inline-block;
		padding: 20rpx 56rpx;
		background: linear-gradient(135deg, #0a84ff, #005ec7);
		color: #fff;
		border-radius: 16rpx;
		font-size: 30rpx;
		font-weight: 600;
		box-shadow: 0 4rpx 16rpx rgba(0, 122, 255, 0.35);
		&:active { transform: scale(0.95); opacity: 0.9; }
	}
}
</style>
