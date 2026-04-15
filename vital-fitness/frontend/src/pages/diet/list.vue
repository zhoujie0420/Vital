<template>
	<view class="page" :style="{ paddingTop: topPadding + 'px' }">
		<view class="page-header">
			<text class="page-title">饮食记录</text>
			<text class="header-action" @tap="addDiet">+ 添加</text>
		</view>

		<!-- 按天分组的卡片 -->
		<view class="day-list">
			<view v-for="(day, i) in groupedDays" :key="i" class="day-card" @tap="goToDetail(day.date)">
				<view class="day-header">
					<view class="day-date-wrap">
						<text class="day-label">{{ day.label }}</text>
						<text class="day-date">{{ day.dateStr }}</text>
					</view>
					<view class="day-summary">
						<text class="day-cal">{{ day.totalCal }}千卡</text>
						<text class="day-arrow">›</text>
					</view>
				</view>
				<view class="day-meals">
					<view v-for="(m, j) in day.meals" :key="j" class="meal-tag" :class="'meal-' + m.type">
						<text>{{ m.name }} {{ m.cal }}kcal</text>
					</view>
				</view>
			</view>
		</view>

		<view v-if="groupedDays.length === 0" class="empty">
			<text class="empty-icon">🥗</text>
			<text class="empty-title">暂无饮食记录</text>
			<text class="empty-desc">记录每一餐，了解你的营养摄入</text>
			<view class="empty-btn" @tap="addDiet"><text>开始记录</text></view>
		</view>
	</view>
</template>

<script>
	import { getDietRecords } from '../../api/diet'

	const mealNames = { breakfast: '早餐', lunch: '午餐', dinner: '晚餐', snack: '加餐' }

	export default {
		data() { return { dietList: [] } },
		computed: {
			topPadding() {
				const app = getApp()
				return (app.globalData?.customBarHeight || 88) + 8
			},
			groupedDays() {
				const map = {}
				this.dietList.forEach(r => {
					const d = new Date(r.record_date)
					const key = d.toISOString().split('T')[0]
					if (!map[key]) map[key] = { date: key, records: [] }
					map[key].records.push(r)
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
					const totalCal = day.records.reduce((s, r) => s + (r.total_calories || 0), 0).toFixed(0)
					const meals = day.records.map(r => ({
						type: r.meal_type,
						name: mealNames[r.meal_type] || r.meal_type,
						cal: (r.total_calories || 0).toFixed(0)
					}))
					return {
						...day, label, totalCal, meals,
						dateStr: d.getFullYear() + '.' + (d.getMonth()+1) + '.' + d.getDate()
					}
				})
			}
		},
		onShow() { this.load() },
		methods: {
			async load() {
				try {
					const res = await getDietRecords({ page: 1, page_size: 200 })
					this.dietList = res.data || []
				} catch (e) {}
			},
			addDiet() { uni.navigateTo({ url: '/pages/diet/add' }) },
			goToDetail(date) {
				uni.navigateTo({ url: '/pages/diet/detail?date=' + date })
			}
		}
	}
</script>

<style lang="scss" scoped>
.page {
	padding: 0 32rpx;
	padding-bottom: 40rpx;
	min-height: 100vh;
	background: #f2f2f7;
}

.page-header {
	display: flex;
	justify-content: space-between;
	align-items: baseline;
	margin-bottom: 28rpx;

	.page-title { font-size: 52rpx; font-weight: 700; color: #1c1c1e; letter-spacing: -1rpx; }
	.header-action { font-size: 30rpx; color: #007aff; font-weight: 600; }
}

.day-card {
	background: #fff;
	border-radius: 20rpx;
	padding: 28rpx 32rpx;
	margin-bottom: 16rpx;
	box-shadow: 0 2rpx 16rpx rgba(0, 0, 0, 0.04);
	transition: transform 0.15s ease;

	&:active { transform: scale(0.98); }

	.day-header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: 16rpx;

		.day-date-wrap {
			.day-label { display: block; font-size: 32rpx; font-weight: 600; color: #1c1c1e; }
			.day-date { display: block; font-size: 22rpx; color: #8e8e93; margin-top: 4rpx; }
		}
		.day-summary {
			display: flex;
			align-items: center;
			gap: 8rpx;

			.day-cal { font-size: 30rpx; font-weight: 700; color: #ff9500; }
			.day-arrow { font-size: 36rpx; color: #c7c7cc; font-weight: 300; }
		}
	}

	.day-meals {
		display: flex;
		flex-wrap: wrap;
		gap: 10rpx;

		.meal-tag {
			padding: 8rpx 20rpx;
			border-radius: 12rpx;
			font-size: 24rpx;
			font-weight: 500;
			background: #f2f2f7;
			color: #8e8e93;
		}
		.meal-breakfast { background: rgba(255, 149, 0, 0.1); color: #ff9500; }
		.meal-lunch { background: rgba(52, 199, 89, 0.1); color: #34c759; }
		.meal-dinner { background: rgba(0, 122, 255, 0.1); color: #007aff; }
		.meal-snack { background: rgba(255, 45, 85, 0.1); color: #ff2d55; }
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
		background: #007aff;
		color: #fff;
		border-radius: 16rpx;
		font-size: 30rpx;
		font-weight: 600;
		box-shadow: 0 4rpx 16rpx rgba(0, 122, 255, 0.3);
		&:active { transform: scale(0.95); opacity: 0.9; }
	}
}
</style>
