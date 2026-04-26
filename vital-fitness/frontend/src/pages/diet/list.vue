<template>
	<view class="page" :style="{ paddingTop: topPadding + 'px' }">
		<view class="page-header">
			<text class="page-title">饮食记录</text>
			<text class="header-action" @tap="addDiet">+ 添加</text>
		</view>

		<!-- 按天分组 -->
		<view class="day-list">
			<view v-for="(day, i) in groupedDays" :key="i" class="day-card" @tap="goToDetail(day.date)">
				<view class="day-header">
					<view class="day-date-wrap">
						<text class="day-label">{{ day.label }}</text>
						<text class="day-date">{{ day.dateStr }}</text>
					</view>
					<view class="day-meta">
						<text class="day-cal">{{ day.totalCal }} 千卡</text>
						<text class="sf-chevron">›</text>
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
		data() { return { dietList: [], page: 1, pageSize: 20, total: 0, loading: false } },
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
		onShow() { this.refresh() },
		onPullDownRefresh() { this.refresh().then(() => uni.stopPullDownRefresh()) },
		onReachBottom() { this.loadMore() },
		methods: {
			async refresh() {
				this.page = 1
				this.dietList = []
				await this.load()
			},
			async loadMore() {
				if (this.loading || this.dietList.length >= this.total) return
				this.page++
				await this.load()
			},
			async load() {
				if (this.loading) return
				this.loading = true
				try {
					const res = await getDietRecords({ page: this.page, page_size: this.pageSize })
					const list = res.data || []
					if (this.page === 1) {
						this.dietList = list
					} else {
						this.dietList = [...this.dietList, ...list]
					}
					this.total = res.total || 0
				} catch (e) {}
				this.loading = false
			},
			addDiet() { uni.navigateTo({ url: '/pages/diet/add' }) },
			goToDetail(date) {
				uni.navigateTo({ url: '/pages/diet/detail?date=' + date })
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
	.header-action { font-size: $font-subhead; color: $color-primary; font-weight: 600; }
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

			.day-cal {
				font-size: $font-subhead;
				font-weight: 700;
				color: $color-orange;
				font-variant-numeric: tabular-nums;
			}
			.sf-chevron {
				font-size: 36rpx;
				color: $color-separator-opaque;
				font-weight: 300;
			}
		}
	}

	.day-meals {
		display: flex;
		flex-wrap: wrap;
		gap: $spacing-xs;

		.meal-tag {
			padding: 6rpx $spacing-md;
			border-radius: $spacing-xs;
			font-size: $font-caption1;
			font-weight: 500;
			background: $color-fill;
			color: $color-label-quaternary;
		}
		.meal-breakfast { background: $color-orange-light; color: $color-orange; }
		.meal-lunch { background: $color-green-light; color: $color-green; }
		.meal-dinner { background: $color-primary-light; color: $color-primary; }
		.meal-snack { background: $color-pink-light; color: $color-pink; }
	}
}

.empty {
	@include empty-state;
}
</style>
