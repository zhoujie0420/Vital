<template>
	<view class="page">
		<view class="page-header">
			<text class="page-title">饮食记录</text>
			<text class="header-action" @tap="addDiet">添加</text>
		</view>

		<view class="list">
			<view v-for="(r, i) in dietList" :key="i" class="list-item">
				<view class="item-top">
					<view class="meal-badge" :class="'meal-' + r.meal_type">
						<text>{{ getMealName(r.meal_type) }}</text>
					</view>
					<text class="item-date">{{ formatDate(r.record_date) }}</text>
				</view>
				<view class="item-main">
					<text class="item-cal">{{ r.total_calories }}</text>
					<text class="item-cal-unit">千卡</text>
				</view>
				<view class="item-macros">
					<view class="macro">
						<text class="macro-val">{{ r.protein }}g</text>
						<text class="macro-label">蛋白质</text>
					</view>
					<view class="macro">
						<text class="macro-val">{{ r.carbs }}g</text>
						<text class="macro-label">碳水</text>
					</view>
					<view class="macro">
						<text class="macro-val">{{ r.fat }}g</text>
						<text class="macro-label">脂肪</text>
					</view>
				</view>
				<text v-if="r.notes" class="item-notes">{{ r.notes }}</text>
			</view>
		</view>

		<view v-if="dietList.length === 0" class="empty">
			<text class="empty-icon">🥗</text>
			<text class="empty-title">暂无饮食记录</text>
			<text class="empty-desc">记录每一餐，了解你的营养摄入</text>
			<view class="empty-btn" @tap="addDiet">
				<text>开始记录</text>
			</view>
		</view>
	</view>
</template>

<script>
	import { getDietRecords } from '../../api/diet'

	export default {
		data() { return { dietList: [] } },
		onShow() { this.load() },
		methods: {
			async load() {
				try {
					const res = await getDietRecords({ page: 1, page_size: 50 })
					this.dietList = res.data || []
				} catch (e) {}
			},
			addDiet() { uni.navigateTo({ url: '/pages/diet/add' }) },
			getMealName(t) {
				return { breakfast: '早餐', lunch: '午餐', dinner: '晚餐', snack: '加餐' }[t] || t
			},
			formatDate(d) {
				if (!d) return ''
				const date = new Date(d)
				return (date.getMonth()+1) + '月' + date.getDate() + '日'
			}
		}
	}
</script>

<style lang="scss" scoped>
.page {
	padding: 0 32rpx;
	padding-top: 120rpx;
	padding-bottom: 40rpx;
	min-height: 100vh;
	background: #f2f2f7;
}

.page-header {
	display: flex;
	justify-content: space-between;
	align-items: baseline;
	margin-bottom: 28rpx;

	.page-title {
		font-size: 52rpx;
		font-weight: 700;
		color: #1c1c1e;
		letter-spacing: -1rpx;
	}
	.header-action {
		font-size: 30rpx;
		color: #007aff;
		font-weight: 500;
	}
}

.list-item {
	background: #fff;
	border-radius: 20rpx;
	padding: 28rpx 32rpx;
	margin-bottom: 16rpx;
	box-shadow: 0 2rpx 16rpx rgba(0, 0, 0, 0.04);

	.item-top {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: 16rpx;
	}

	.meal-badge {
		padding: 6rpx 20rpx;
		border-radius: 12rpx;
		font-size: 24rpx;
		font-weight: 600;
		background: #f2f2f7;
		color: #8e8e93;
	}
	.meal-breakfast { background: #fff3e0; color: #ff9500; }
	.meal-lunch { background: #e8f5e9; color: #34c759; }
	.meal-dinner { background: #e3f2fd; color: #007aff; }
	.meal-snack { background: #fce4ec; color: #ff2d55; }

	.item-date {
		font-size: 26rpx;
		color: #8e8e93;
	}

	.item-main {
		display: flex;
		align-items: baseline;
		margin-bottom: 20rpx;

		.item-cal {
			font-size: 56rpx;
			font-weight: 700;
			color: #1c1c1e;
			letter-spacing: -2rpx;
		}
		.item-cal-unit {
			font-size: 26rpx;
			color: #8e8e93;
			margin-left: 8rpx;
		}
	}

	.item-macros {
		display: flex;
		gap: 32rpx;

		.macro {
			.macro-val {
				display: block;
				font-size: 28rpx;
				font-weight: 600;
				color: #3a3a3c;
			}
			.macro-label {
				display: block;
				font-size: 22rpx;
				color: #8e8e93;
				margin-top: 2rpx;
			}
		}
	}

	.item-notes {
		display: block;
		font-size: 26rpx;
		color: #8e8e93;
		margin-top: 16rpx;
		line-height: 1.5;
	}
}

.empty {
	text-align: center;
	padding: 120rpx 0;

	.empty-icon {
		display: block;
		font-size: 96rpx;
		margin-bottom: 24rpx;
	}
	.empty-title {
		display: block;
		font-size: 34rpx;
		font-weight: 600;
		color: #1c1c1e;
		margin-bottom: 8rpx;
	}
	.empty-desc {
		display: block;
		font-size: 26rpx;
		color: #8e8e93;
		margin-bottom: 40rpx;
	}
	.empty-btn {
		display: inline-block;
		padding: 20rpx 56rpx;
		background: #007aff;
		color: #fff;
		border-radius: 16rpx;
		font-size: 30rpx;
		font-weight: 600;

		&:active { opacity: 0.8; }
	}
}
</style>
