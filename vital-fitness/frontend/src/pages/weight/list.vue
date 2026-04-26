<template>
	<view class="page" :style="{ paddingTop: topPadding + 'px' }">
		<view class="page-header">
			<text class="nav-back" @tap="goBack">‹</text>
			<text class="page-title">体重记录</text>
			<text class="header-action" @tap="addWeight">+ 添加</text>
		</view>

		<view class="list">
			<view v-for="(r, i) in weightList" :key="i" class="weight-card">
				<view class="weight-row">
					<view class="weight-left">
						<view class="weight-main">
							<text class="weight-val">{{ r.weight }}</text>
							<text class="weight-unit">kg</text>
						</view>
						<text v-if="r.bmi" class="weight-bmi">BMI {{ r.bmi }}</text>
					</view>
					<view class="weight-right">
						<text class="weight-date">{{ formatDate(r.record_date) }}</text>
						<view v-if="i < weightList.length - 1" class="weight-trend">
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
		data() { return { weightList: [], page: 1, pageSize: 20, total: 0, loading: false } },
		computed: {
			topPadding() {
				const app = getApp()
				return (app.globalData?.customBarHeight || 88) + 8
			}
		},
		onShow() { this.refresh() },
		onPullDownRefresh() { this.refresh().then(() => uni.stopPullDownRefresh()) },
		onReachBottom() { this.loadMore() },
		methods: {
			async refresh() {
				this.page = 1
				this.weightList = []
				await this.load()
			},
			async loadMore() {
				if (this.loading || this.weightList.length >= this.total) return
				this.page++
				await this.load()
			},
			async load() {
				if (this.loading) return
				this.loading = true
				try {
					const res = await getWeights({ page: this.page, page_size: this.pageSize })
					const list = res.data || []
					if (this.page === 1) {
						this.weightList = list
					} else {
						this.weightList = [...this.weightList, ...list]
					}
					this.total = res.total || 0
				} catch (e) {}
				this.loading = false
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
	align-items: center;
	margin-bottom: $spacing-lg;

	.nav-back {
		font-size: 48rpx;
		color: $color-primary;
		font-weight: 300;
		line-height: 1;
		width: 60rpx;
	}
	.page-title {
		font-size: $font-headline;
		font-weight: 600;
		color: $color-label;
	}
	.header-action {
		font-size: $font-subhead;
		color: $color-primary;
		font-weight: 600;
		width: 60rpx;
		text-align: right;
	}
}

// --- Weight Card ---
.weight-card {
	@include card;
	margin-bottom: $spacing-sm;
	@include press-effect;

	.weight-row {
		display: flex;
		justify-content: space-between;
		align-items: center;
	}

	.weight-left {
		.weight-main {
			display: flex;
			align-items: baseline;

			.weight-val {
				font-size: 52rpx;
				font-weight: 700;
				color: $color-label;
				letter-spacing: -2rpx;
				font-variant-numeric: tabular-nums;
			}
			.weight-unit {
				font-size: $font-footnote;
				color: $color-label-quaternary;
				margin-left: 6rpx;
			}
		}
		.weight-bmi {
			display: block;
			font-size: $font-caption1;
			color: $color-label-quaternary;
			margin-top: 2rpx;
		}
	}

	.weight-right {
		text-align: right;

		.weight-date {
			display: block;
			font-size: $font-footnote;
			color: $color-label-quaternary;
		}
		.weight-trend {
			margin-top: 6rpx;
			font-size: $font-footnote;
			font-weight: 600;
			font-variant-numeric: tabular-nums;

			.trend-up { color: $color-red; }
			.trend-down { color: $color-green; }
			.trend-flat { color: $color-label-quaternary; }
		}
	}
}

.empty {
	@include empty-state;
}
</style>
