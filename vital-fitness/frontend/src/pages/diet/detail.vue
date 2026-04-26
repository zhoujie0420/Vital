<template>
	<view class="page" :style="{ paddingTop: topPadding + 'px' }">
		<view class="nav-bar">
			<text class="nav-back" @tap="goBack">‹ 返回</text>
			<text class="nav-title">{{ dayLabel }}</text>
			<text class="nav-action" @tap="goAdd">+ 添加</text>
		</view>

		<!-- 当天汇总 -->
		<view class="summary-card" v-if="records.length > 0">
			<view class="summary-main">
				<text class="summary-val">{{ totalCal }}</text>
				<text class="summary-unit">千卡</text>
			</view>
			<view class="summary-macros">
				<view class="macro">
					<view class="macro-dot dot-blue"></view>
					<text class="macro-text">蛋白 {{ totalProtein }}g</text>
				</view>
				<view class="macro">
					<view class="macro-dot dot-orange"></view>
					<text class="macro-text">碳水 {{ totalCarbs }}g</text>
				</view>
				<view class="macro">
					<view class="macro-dot dot-yellow"></view>
					<text class="macro-text">脂肪 {{ totalFat }}g</text>
				</view>
			</view>
		</view>

		<!-- 每餐记录 -->
		<view v-for="(r, i) in records" :key="r.id" class="meal-card">
			<view class="meal-header">
				<view class="meal-badge" :class="'badge-' + r.meal_type">
					<text>{{ getMealName(r.meal_type) }}</text>
				</view>
				<text class="meal-cal">{{ r.total_calories }} kcal</text>
			</view>
			<view class="meal-macros">
				<text class="meal-macro">蛋白 {{ r.protein }}g</text>
				<text class="meal-sep">·</text>
				<text class="meal-macro">碳水 {{ r.carbs }}g</text>
				<text class="meal-sep">·</text>
				<text class="meal-macro">脂肪 {{ r.fat }}g</text>
			</view>
			<text v-if="r.notes" class="meal-notes">{{ r.notes }}</text>
			<view class="meal-actions">
				<text class="action-link action-delete" @tap="confirmDelete(r)">删除</text>
			</view>
		</view>

		<view v-if="records.length === 0" class="empty">
			<text class="empty-icon">🍽️</text>
			<text class="empty-title">当天暂无饮食记录</text>
			<text class="empty-desc">点击右上角添加记录</text>
		</view>
	</view>
</template>

<script>
	import { getDietRecords, deleteDietRecord } from '../../api/diet'

	export default {
		data() { return { date: '', records: [] } },
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
			totalProtein() { return this.records.reduce((s, r) => s + (r.protein || 0), 0).toFixed(1) },
			totalCarbs() { return this.records.reduce((s, r) => s + (r.carbs || 0), 0).toFixed(1) },
			totalFat() { return this.records.reduce((s, r) => s + (r.fat || 0), 0).toFixed(1) }
		},
		onLoad(opts) { this.date = opts.date || '' },
		onShow() { this.load() },
		methods: {
			goBack() { uni.navigateBack() },
			goAdd() { uni.navigateTo({ url: '/pages/diet/add' }) },
			getMealName(t) {
				return { breakfast: '早餐', lunch: '午餐', dinner: '晚餐', snack: '加餐' }[t] || t
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
				uni.showModal({
					title: '确认删除',
					content: `删除这条${this.getMealName(r.meal_type)}记录？`,
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
	border-radius: $radius-xl;
	padding: $spacing-xl;
	margin-bottom: $spacing-md;
	background: linear-gradient(145deg, #1C1C1E 0%, #2C2C2E 100%);

	.summary-main {
		display: flex;
		align-items: baseline;
		margin-bottom: $spacing-md;

		.summary-val {
			font-size: 60rpx;
			font-weight: 700;
			color: #fff;
			letter-spacing: -2rpx;
			font-variant-numeric: tabular-nums;
		}
		.summary-unit {
			font-size: $font-footnote;
			color: rgba(255, 255, 255, 0.5);
			margin-left: $spacing-xs;
		}
	}

	.summary-macros {
		display: flex;
		gap: $spacing-lg;

		.macro {
			display: flex;
			align-items: center;
			gap: $spacing-xs;

			.macro-dot {
				width: 12rpx;
				height: 12rpx;
				border-radius: 50%;
			}
			.dot-blue { background: $color-teal; }
			.dot-orange { background: $color-orange; }
			.dot-yellow { background: $color-yellow; }
			.macro-text {
				font-size: $font-caption1;
				color: rgba(255, 255, 255, 0.65);
				font-weight: 500;
			}
		}
	}
}

// --- Meal Card ---
.meal-card {
	@include card;
	margin-bottom: $spacing-sm;

	.meal-header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: $spacing-md;

		.meal-badge {
			padding: 6rpx $spacing-md;
			border-radius: $spacing-xs;
			font-size: $font-caption1;
			font-weight: 600;
		}
		.badge-breakfast { background: $color-orange-light; color: $color-orange; }
		.badge-lunch { background: $color-green-light; color: $color-green; }
		.badge-dinner { background: $color-primary-light; color: $color-primary; }
		.badge-snack { background: $color-pink-light; color: $color-pink; }

		.meal-cal {
			font-size: $font-headline;
			font-weight: 700;
			color: $color-label;
			font-variant-numeric: tabular-nums;
		}
	}

	.meal-macros {
		display: flex;
		align-items: center;
		gap: $spacing-xs;
		margin-bottom: $spacing-sm;

		.meal-macro {
			font-size: $font-caption1;
			color: $color-label-quaternary;
			font-weight: 500;
		}
		.meal-sep {
			font-size: $font-caption1;
			color: $color-separator-opaque;
		}
	}

	.meal-notes {
		display: block;
		font-size: $font-footnote;
		color: $color-label-quaternary;
		margin-bottom: $spacing-sm;
		line-height: 1.5;
	}

	.meal-actions {
		padding-top: $spacing-md;
		border-top: 0.5rpx solid $color-separator;

		.action-link {
			font-size: $font-subhead;
			font-weight: 500;
			padding: $spacing-xs 0;

			&.action-delete { color: $color-red; }
			&:active { opacity: 0.5; }
		}
	}
}

.empty {
	@include empty-state;
}
</style>
