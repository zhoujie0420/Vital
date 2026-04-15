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
				<text class="meal-cal">{{ r.total_calories }}kcal</text>
			</view>
			<view class="meal-macros">
				<text class="meal-macro">蛋白{{ r.protein }}g</text>
				<text class="meal-macro">碳水{{ r.carbs }}g</text>
				<text class="meal-macro">脂肪{{ r.fat }}g</text>
			</view>
			<text v-if="r.notes" class="meal-notes">{{ r.notes }}</text>
			<view class="meal-actions">
				<text class="action-btn delete" @tap="confirmDelete(r)">删除</text>
			</view>
		</view>

		<view v-if="records.length === 0" class="empty">
			<text class="empty-text">当天暂无饮食记录</text>
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
.page {
	padding: 0 32rpx;
	padding-bottom: 40rpx;
	min-height: 100vh;
	background: #f2f2f7;
}

.nav-bar {
	display: flex;
	align-items: center;
	justify-content: space-between;
	margin-bottom: 28rpx;

	.nav-back { font-size: 32rpx; color: #007aff; font-weight: 500; }
	.nav-title { font-size: 34rpx; font-weight: 600; color: #1c1c1e; }
	.nav-action { font-size: 30rpx; color: #007aff; font-weight: 600; }
}

.summary-card {
	background: linear-gradient(135deg, #ff9500, #ff6b00);
	border-radius: 20rpx;
	padding: 32rpx;
	margin-bottom: 20rpx;

	.summary-main {
		display: flex;
		align-items: baseline;
		margin-bottom: 16rpx;

		.summary-val { font-size: 56rpx; font-weight: 700; color: #fff; letter-spacing: -2rpx; }
		.summary-unit { font-size: 26rpx; color: rgba(255,255,255,0.8); margin-left: 8rpx; }
	}

	.summary-macros {
		display: flex;
		gap: 24rpx;

		.macro {
			display: flex;
			align-items: center;
			gap: 8rpx;

			.macro-dot { width: 12rpx; height: 12rpx; border-radius: 50%; }
			.dot-blue { background: rgba(255,255,255,0.9); }
			.dot-orange { background: rgba(255,255,255,0.7); }
			.dot-yellow { background: rgba(255,255,255,0.5); }
			.macro-text { font-size: 24rpx; color: rgba(255,255,255,0.85); font-weight: 500; }
		}
	}
}

.meal-card {
	background: #fff;
	border-radius: 20rpx;
	padding: 28rpx 32rpx;
	margin-bottom: 16rpx;
	box-shadow: 0 2rpx 16rpx rgba(0, 0, 0, 0.04);

	.meal-header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: 16rpx;

		.meal-badge {
			padding: 6rpx 20rpx;
			border-radius: 12rpx;
			font-size: 24rpx;
			font-weight: 600;
		}
		.badge-breakfast { background: rgba(255, 149, 0, 0.12); color: #ff9500; }
		.badge-lunch { background: rgba(52, 199, 89, 0.12); color: #34c759; }
		.badge-dinner { background: rgba(0, 122, 255, 0.12); color: #007aff; }
		.badge-snack { background: rgba(255, 45, 85, 0.12); color: #ff2d55; }

		.meal-cal { font-size: 32rpx; font-weight: 700; color: #1c1c1e; }
	}

	.meal-macros {
		display: flex;
		gap: 20rpx;
		margin-bottom: 12rpx;

		.meal-macro { font-size: 24rpx; color: #8e8e93; font-weight: 500; }
	}

	.meal-notes {
		display: block;
		font-size: 26rpx;
		color: #8e8e93;
		margin-bottom: 12rpx;
		line-height: 1.5;
	}

	.meal-actions {
		padding-top: 16rpx;
		border-top: 1rpx solid #f2f2f7;

		.action-btn {
			font-size: 28rpx;
			font-weight: 500;
			padding: 8rpx 0;

			&.delete { color: #ff3b30; }
			&:active { opacity: 0.6; }
		}
	}
}

.empty {
	text-align: center;
	padding: 80rpx 0;
	.empty-text { font-size: 28rpx; color: #8e8e93; }
}
</style>
