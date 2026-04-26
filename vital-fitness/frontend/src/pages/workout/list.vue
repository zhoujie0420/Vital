<template>
	<view class="page" :style="{ paddingTop: topPadding + 'px' }">
		<view class="page-header">
			<text class="page-title">训练记录</text>
			<text class="header-action" @tap="addWorkout">+ 记录</text>
		</view>

		<!-- 分类筛选 -->
		<scroll-view scroll-x class="filter-scroll" :show-scrollbar="false">
			<view class="filter-bar">
				<text v-for="(tab, index) in tabList" :key="index" class="filter-pill"
					:class="{ active: currentTab === index }" @tap="changeTab(index)">
					{{ tab.name }}
				</text>
			</view>
		</scroll-view>

		<!-- 按天分组 -->
		<view class="day-list">
			<view v-for="(day, i) in groupedDays" :key="i" class="day-card" @tap="goToDetail(day.date)">
				<view class="day-header">
					<view class="day-date-wrap">
						<text class="day-label">{{ day.label }}</text>
						<text class="day-date">{{ day.dateStr }}</text>
					</view>
					<view class="day-meta">
						<text class="day-count">{{ day.workouts.length }} 个动作</text>
						<text class="sf-chevron">›</text>
					</view>
				</view>
				<view class="day-tags">
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
				workoutList: [],
				page: 1,
				pageSize: 20,
				total: 0,
				loading: false
			}
		},
		onShow() { this.refresh() },
		onPullDownRefresh() { this.refresh().then(() => uni.stopPullDownRefresh()) },
		onReachBottom() { this.loadMore() },
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
				this.refresh()
			},
			async refresh() {
				this.page = 1
				this.workoutList = []
				await this.loadWorkouts()
			},
			async loadMore() {
				if (this.loading || this.workoutList.length >= this.total) return
				this.page++
				await this.loadWorkouts()
			},
			async loadWorkouts() {
				if (this.loading) return
				this.loading = true
				try {
					const category = this.tabList[this.currentTab].name
					const res = await getWorkouts({
						category: category === '全部' ? '' : category,
						page: this.page, page_size: this.pageSize
					})
					const list = res.data || []
					if (this.page === 1) {
						this.workoutList = list
					} else {
						this.workoutList = [...this.workoutList, ...list]
					}
					this.total = res.total || 0
				} catch (e) {}
				this.loading = false
			},
			addWorkout() { uni.navigateTo({ url: '/pages/workout/add' }) },
			goToDetail(date) {
				uni.navigateTo({ url: '/pages/workout/detail?date=' + date })
			}
		}
	}
</script>

<style lang="scss" scoped>
@import '../../styles/variables.scss';

.page {
	padding: 0 $spacing-xl;
	padding-bottom: 40rpx;
	min-height: 100vh;
	background: $color-bg;
}

.page-header {
	display: flex;
	justify-content: space-between;
	align-items: baseline;
	margin-bottom: $spacing-lg;

	.page-title { @include page-title; }
	.header-action {
		font-size: $font-subhead;
		color: $color-primary;
		font-weight: 600;
	}
}

// --- Filter ---
.filter-scroll {
	white-space: nowrap;
	margin: 0 -#{$spacing-xl};
	padding: 0 $spacing-xl;
	margin-bottom: $spacing-lg;

	.filter-bar {
		display: inline-flex;
		gap: $spacing-sm;

		.filter-pill {
			display: inline-block;
			padding: $spacing-sm $spacing-lg;
			border-radius: $radius-full;
			background: $color-fill;
			font-size: $font-footnote;
			color: $color-label-tertiary;
			font-weight: 500;
			transition: all 0.25s cubic-bezier(0.25, 0.46, 0.45, 0.94);

			&.active {
				background: $color-label;
				color: #fff;
			}
			&:active { transform: scale(0.94); }
		}
	}
}

// --- Day Card ---
.day-card {
	@include card;
	margin-bottom: $spacing-sm;
	@include press-effect;

	.day-header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: $spacing-md;

		.day-date-wrap {
			.day-label {
				display: block;
				font-size: $font-headline;
				font-weight: 600;
				color: $color-label;
			}
			.day-date {
				display: block;
				font-size: $font-caption2;
				color: $color-label-quaternary;
				margin-top: 4rpx;
			}
		}

		.day-meta {
			display: flex;
			align-items: center;
			gap: $spacing-xs;

			.day-count {
				font-size: $font-footnote;
				color: $color-label-quaternary;
				font-weight: 500;
			}
			.sf-chevron {
				font-size: 36rpx;
				color: $color-separator-opaque;
				font-weight: 300;
			}
		}
	}

	.day-tags {
		display: flex;
		flex-wrap: wrap;
		gap: $spacing-xs;

		.exercise-tag {
			padding: 6rpx $spacing-md;
			background: $color-primary-light;
			color: $color-primary;
			border-radius: $spacing-xs;
			font-size: $font-caption1;
			font-weight: 500;
		}
		.exercise-more {
			padding: 6rpx $spacing-sm;
			color: $color-label-quaternary;
			font-size: $font-caption1;
			font-weight: 500;
		}
	}
}

// --- Empty ---
.empty {
	@include empty-state;
}
</style>
