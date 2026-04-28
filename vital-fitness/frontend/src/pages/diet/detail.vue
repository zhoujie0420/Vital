<template>
	<view class="page" :style="{ paddingTop: topPadding + 'px' }">
		<view class="nav-bar">
			<text class="nav-back" @tap="goBack">‹ 返回</text>
			<text class="nav-title">{{ dayLabel }}</text>
			<text class="nav-action" @tap="goAdd">+ 添加</text>
		</view>

		<!-- 当天汇总 -->
		<view class="summary-card" v-if="records.length > 0">
			<view class="summary-top">
				<text class="summary-cal">{{ totalCal }}</text>
				<text class="summary-unit">千卡</text>
			</view>
			<view class="summary-macros">
				<view class="macro-col">
					<view class="macro-track"><view class="macro-fill fill-teal" :style="{ width: macroPercent(totalProtein, 120) }"></view></view>
					<text class="macro-label">蛋白 {{ totalProtein }}g</text>
				</view>
				<view class="macro-col">
					<view class="macro-track"><view class="macro-fill fill-orange" :style="{ width: macroPercent(totalCarbs, 250) }"></view></view>
					<text class="macro-label">碳水 {{ totalCarbs }}g</text>
				</view>
				<view class="macro-col">
					<view class="macro-track"><view class="macro-fill fill-yellow" :style="{ width: macroPercent(totalFat, 65) }"></view></view>
					<text class="macro-label">脂肪 {{ totalFat }}g</text>
				</view>
			</view>
		</view>

		<!-- 按餐次分组 -->
		<view v-for="meal in mealOrder" :key="meal.type">
			<view v-if="getMealRecords(meal.type).length > 0" class="meal-group">
				<view class="meal-group-header">
					<view class="meal-group-left">
						<view class="meal-dot" :class="'dot-' + meal.type"></view>
						<text class="meal-group-title">{{ meal.name }}</text>
					</view>
					<text class="meal-group-cal">{{ getMealCal(meal.type) }} kcal</text>
				</view>

				<view v-for="r in getMealRecords(meal.type)" :key="r.id" class="meal-record">
					<view class="meal-record-info">
						<text class="meal-record-cal">{{ r.total_calories }} kcal</text>
						<text class="meal-record-macro">蛋白{{ r.protein }}g · 碳水{{ r.carbs }}g · 脂肪{{ r.fat }}g</text>
						<text v-if="r.notes" class="meal-record-notes">{{ r.notes }}</text>
					</view>
					<text class="meal-record-delete" @tap="confirmDelete(r)">删除</text>
				</view>
			</view>
		</view>

		<view v-if="records.length === 0" class="empty">
			<view class="empty-icon">
				<text class="empty-icon-letter">D</text>
			</view>
			<text class="empty-title">当天暂无饮食记录</text>
			<text class="empty-desc">点击右上角添加记录</text>
		</view>
	</view>
</template>

<script>
	import { getDietRecords, deleteDietRecord } from '../../api/diet'

	export default {
		data() {
			return {
				date: '', records: [],
				mealOrder: [
					{ type: 'breakfast', name: '早餐' },
					{ type: 'lunch', name: '午餐' },
					{ type: 'dinner', name: '晚餐' },
					{ type: 'snack', name: '加餐' }
				]
			}
		},
		computed: {
			topPadding() {
				const app = getApp()
				return (app.globalData?.customBarHeight || 88) + 8
			},
			dayLabel() {
				if (!this.date) return ''
				const d = new Date(this.date + 'T00:00:00')
				const today = new Date()
				const yesterday = new Date(today)
				yesterday.setDate(yesterday.getDate() - 1)
				if (d.toDateString() === today.toDateString()) return '今天的饮食'
				if (d.toDateString() === yesterday.toDateString()) return '昨天的饮食'
				return (d.getMonth() + 1) + '月' + d.getDate() + '日饮食'
			},
			totalCal() { return this.records.reduce((s, r) => s + (r.total_calories || 0), 0).toFixed(0) },
			totalProtein() { return this.records.reduce((s, r) => s + (r.protein || 0), 0).toFixed(0) },
			totalCarbs() { return this.records.reduce((s, r) => s + (r.carbs || 0), 0).toFixed(0) },
			totalFat() { return this.records.reduce((s, r) => s + (r.fat || 0), 0).toFixed(0) }
		},
		onLoad(opts) { this.date = opts.date || '' },
		onShow() { this.load() },
		methods: {
			goBack() { uni.navigateBack() },
			goAdd() { uni.navigateTo({ url: '/pages/diet/add' }) },
			getMealRecords(type) { return this.records.filter(r => r.meal_type === type) },
			getMealCal(type) { return this.getMealRecords(type).reduce((s, r) => s + (r.total_calories || 0), 0).toFixed(0) },
			macroPercent(val, target) {
				return Math.min((parseFloat(val) / target) * 100, 100) + '%'
			},
			async load() {
				if (!this.date) return
				try {
					const res = await getDietRecords({ page: 1, page_size: 200 })
					const all = res.data || []
					this.records = all.filter(r => {
						const d = new Date(r.record_date)
						return d.toISOString().split('T')[0] === this.date
					})
				} catch (e) {}
			},
			confirmDelete(r) {
				const name = { breakfast: '早餐', lunch: '午餐', dinner: '晚餐', snack: '加餐' }[r.meal_type] || ''
				uni.showModal({
					title: '确认删除',
					content: `删除这条${name}记录？`,
					success: async (res) => {
						if (res.confirm) {
							try {
								await deleteDietRecord(r.id)
								uni.showToast({ title: '已删除', icon: 'success' })
								this.load()
							} catch (e) {}
						}
					}
				})
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

.nav-bar {
	@include nav-bar;
	.nav-back { font-size: $font-callout; color: $color-primary; font-weight: 500; }
	.nav-title { font-size: $font-headline; font-weight: 600; color: $color-label; }
	.nav-action { font-size: $font-subhead; color: $color-primary; font-weight: 600; }
}

// --- Summary Card ---
.summary-card {
	@include card;
	padding: $spacing-xl;
	margin-bottom: $spacing-lg;

	.summary-top {
		display: flex;
		align-items: baseline;
		margin-bottom: $spacing-lg;

		.summary-cal {
			font-size: 56rpx;
			font-weight: 700;
			color: $color-primary;
			letter-spacing: -2rpx;
			font-variant-numeric: tabular-nums;
		}
		.summary-unit {
			font-size: $font-footnote;
			color: $color-label-quaternary;
			margin-left: $spacing-xs;
		}
	}

	.summary-macros {
		display: flex;
		gap: $spacing-md;
	}
}

.macro-col {
	flex: 1;

	.macro-label {
		display: block;
		font-size: $font-caption2;
		color: $color-label-tertiary;
		font-weight: 500;
		margin-top: $spacing-xs;
		text-align: center;
	}
}

.macro-track {
	height: 12rpx;
	background: $color-fill;
	border-radius: 6rpx;
	overflow: hidden;
}

.macro-fill {
	height: 100%;
	border-radius: 6rpx;
	min-width: 4rpx;
	transition: width 0.4s cubic-bezier(0.16, 1, 0.3, 1);
}
.fill-teal { background: $color-teal; }
.fill-orange { background: $color-orange; }
.fill-yellow { background: $color-yellow; }

// --- Meal Group ---
.meal-group {
	@include card;
	padding: $spacing-lg $spacing-xl;
	margin-bottom: $spacing-sm;
}

.meal-group-header {
	display: flex;
	justify-content: space-between;
	align-items: center;
}

.meal-group-left {
	display: flex;
	align-items: center;
	gap: $spacing-sm;
}

.meal-dot {
	width: 14rpx;
	height: 14rpx;
	border-radius: 50%;
}
.dot-breakfast { background: $color-orange; }
.dot-lunch { background: $color-green; }
.dot-dinner { background: $color-primary; }
.dot-snack { background: $color-pink; }

.meal-group-title {
	font-size: $font-headline;
	font-weight: 600;
	color: $color-label;
}

.meal-group-cal {
	font-size: $font-subhead;
	font-weight: 700;
	color: $color-accent;
	font-variant-numeric: tabular-nums;
}

.meal-record {
	display: flex;
	align-items: center;
	justify-content: space-between;
	padding: $spacing-md 0;
	border-top: 0.5rpx solid $color-separator;
	margin-top: $spacing-md;

	.meal-record-info {
		flex: 1;

		.meal-record-cal {
			display: block;
			font-size: $font-body;
			font-weight: 600;
			color: $color-label;
			font-variant-numeric: tabular-nums;
		}
		.meal-record-macro {
			display: block;
			font-size: $font-caption2;
			color: $color-label-quaternary;
			margin-top: 4rpx;
		}
		.meal-record-notes {
			display: block;
			font-size: $font-caption1;
			color: $color-label-tertiary;
			margin-top: 4rpx;
			line-height: 1.4;
		}
	}

	.meal-record-delete {
		font-size: $font-caption1;
		color: $color-red;
		font-weight: 500;
		padding: $spacing-xs $spacing-sm;
		&:active { opacity: 0.5; }
	}
}

.empty {
	@include empty-state;

	.empty-icon-letter {
		font-size: 44rpx;
		font-weight: 700;
		color: $color-label-quaternary;
	}
}
</style>
