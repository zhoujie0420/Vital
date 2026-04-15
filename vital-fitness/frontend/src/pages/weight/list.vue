<template>
	<view class="page" :style="{ paddingTop: topPadding + 'px' }">
		<view class="page-header">
			<text class="nav-back" @tap="goBack">‹</text>
			<text class="page-title">体重记录</text>
			<text class="header-action" @tap="addWeight">+ 添加</text>
		</view>

		<view class="list">
			<view v-for="(r, i) in weightList" :key="i" class="list-item">
				<view class="item-row">
					<view class="item-left">
						<view class="item-main">
							<text class="item-weight">{{ r.weight }}</text>
							<text class="item-unit">kg</text>
						</view>
						<text v-if="r.bmi" class="item-bmi">BMI {{ r.bmi }}</text>
					</view>
					<view class="item-right">
						<text class="item-date">{{ formatDate(r.record_date) }}</text>
						<view v-if="i < weightList.length - 1" class="item-trend">
							<text :class="getTrend(r, weightList[i+1]).cls">{{ getTrend(r, weightList[i+1]).text }}</text>
						</view>
					</view>
				</view>
			</view>
		</view>

		<view v-if="weightList.length === 0" class="empty">
			<text class="empty-icon">⚖️</text>
			<text class="empty-title">暂无体重记录</text>
			<text class="empty-desc">定期记录体重，追踪你的变化</text>
			<view class="empty-btn" @tap="addWeight">
				<text>开始记录</text>
			</view>
		</view>
	</view>
</template>

<script>
	import { getWeights } from '../../api/weight'

	export default {
		data() { return { weightList: [] } },
		computed: {
			topPadding() {
				const app = getApp()
				return (app.globalData?.customBarHeight || 88) + 8
			}
		},
		onShow() { this.load() },
		methods: {
			async load() {
				try {
					const res = await getWeights({ page: 1, page_size: 50 })
					this.weightList = res.data || []
				} catch (e) {}
			},
			addWeight() { uni.navigateTo({ url: '/pages/weight/add' }) },
			goBack() { uni.navigateBack() },
			formatDate(d) {
				if (!d) return ''
				const date = new Date(d)
				return (date.getMonth()+1) + '月' + date.getDate() + '日'
			},
			getTrend(curr, prev) {
				const diff = (curr.weight - prev.weight).toFixed(1)
				if (diff > 0) return { text: '+' + diff, cls: 'trend-up' }
				if (diff < 0) return { text: diff, cls: 'trend-down' }
				return { text: '持平', cls: 'trend-flat' }
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
	align-items: center;
	margin-bottom: 28rpx;

	.nav-back {
		font-size: 48rpx;
		color: #007aff;
		font-weight: 300;
		line-height: 1;
	}

	.page-title {
		font-size: 34rpx;
		font-weight: 600;
		color: #1c1c1e;
	}
	.header-action {
		font-size: 30rpx;
		color: #007aff;
		font-weight: 600;
	}
}

.list-item {
	background: #fff;
	border-radius: 20rpx;
	padding: 28rpx 32rpx;
	margin-bottom: 16rpx;
	box-shadow: 0 2rpx 16rpx rgba(0, 0, 0, 0.04);
	transition: transform 0.15s ease;

	&:active { transform: scale(0.98); }

	.item-row {
		display: flex;
		justify-content: space-between;
		align-items: center;
	}

	.item-left {
		.item-main {
			display: flex;
			align-items: baseline;

			.item-weight {
				font-size: 56rpx;
				font-weight: 700;
				color: #1c1c1e;
				letter-spacing: -2rpx;
			}
			.item-unit {
				font-size: 26rpx;
				color: #8e8e93;
				margin-left: 6rpx;
			}
		}
		.item-bmi {
			display: block;
			font-size: 24rpx;
			color: #8e8e93;
			margin-top: 4rpx;
		}
	}

	.item-right {
		text-align: right;

		.item-date {
			display: block;
			font-size: 26rpx;
			color: #8e8e93;
		}
		.item-trend {
			margin-top: 6rpx;
			font-size: 26rpx;
			font-weight: 600;

			.trend-up { color: #ff3b30; }
			.trend-down { color: #34c759; }
			.trend-flat { color: #8e8e93; }
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
