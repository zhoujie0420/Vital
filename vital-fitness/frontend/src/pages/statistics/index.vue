<template>
	<view class="page" :style="{ paddingTop: topPadding + 'px' }">
		<view class="page-header">
			<text class="page-title">数据统计</text>
		</view>

		<!-- 时间段选择 -->
		<view class="period-bar">
			<text v-for="p in periods" :key="p.key" class="period-item"
				:class="{ active: currentPeriod === p.key }" @tap="currentPeriod = p.key">
				{{ p.label }}
			</text>
		</view>

		<!-- 训练统计 -->
		<view class="stat-card">
			<view class="stat-header">
				<view class="stat-header-left">
					<view class="stat-dot dot-blue"></view>
					<text class="stat-label">训练次数</text>
				</view>
				<text class="stat-icon">💪</text>
			</view>
			<view v-if="stats.workout_stats && stats.workout_stats.length > 0" class="chart-area">
				<view class="bar-chart">
					<view v-for="(s, i) in stats.workout_stats" :key="i" class="bar-col">
						<text class="bar-val" v-if="s.count > 0">{{ s.count }}</text>
						<view class="bar-track">
							<view class="bar-fill fill-blue" :style="{ height: getWorkoutBarHeight(s.count) }"></view>
						</view>
						<text class="bar-day">{{ formatDay(s.day) }}</text>
					</view>
				</view>
			</view>
			<view v-else class="no-data">
				<text class="no-data-text">暂无训练数据</text>
			</view>
		</view>

		<!-- 体重趋势 -->
		<view class="stat-card">
			<view class="stat-header">
				<view class="stat-header-left">
					<view class="stat-dot dot-green"></view>
					<text class="stat-label">体重趋势</text>
				</view>
				<text class="stat-icon">⚖️</text>
			</view>
			<view v-if="stats.weight_trend && stats.weight_trend.length > 0" class="trend-area">
				<view v-for="(w, i) in stats.weight_trend" :key="i" class="trend-row">
					<text class="trend-day">{{ formatDay(w.day) }}</text>
					<view class="trend-bar-track">
						<view class="trend-bar-fill" :style="{ width: getBarWidth(w.weight) + '%' }"></view>
					</view>
					<text class="trend-val">{{ w.weight }}<text class="trend-unit">kg</text></text>
				</view>
			</view>
			<view v-else class="no-data">
				<text class="no-data-text">暂无体重数据</text>
			</view>
		</view>

		<!-- 饮食热量 -->
		<view class="stat-card">
			<view class="stat-header">
				<view class="stat-header-left">
					<view class="stat-dot dot-orange"></view>
					<text class="stat-label">热量摄入</text>
				</view>
				<text class="stat-icon">🍽️</text>
			</view>
			<view v-if="stats.diet_stats && stats.diet_stats.length > 0" class="chart-area">
				<view class="bar-chart">
					<view v-for="(s, i) in stats.diet_stats" :key="i" class="bar-col">
						<text class="bar-val" v-if="s.calories > 0">{{ s.calories }}</text>
						<view class="bar-track">
							<view class="bar-fill fill-orange" :style="{ height: getDietBarHeight(s.calories) }"></view>
						</view>
						<text class="bar-day">{{ formatDay(s.day) }}</text>
					</view>
				</view>
			</view>
			<view v-else class="no-data">
				<text class="no-data-text">暂无饮食数据</text>
			</view>
		</view>
	</view>
</template>

<script>
	import { getStats } from '../../api/stats'

	export default {
		data() {
			return {
				stats: {},
				currentPeriod: 'week',
				periods: [
					{ key: 'week', label: '周' },
					{ key: 'month', label: '月' },
					{ key: 'year', label: '年' }
				]
			}
		},
		computed: {
			topPadding() {
				const app = getApp()
				return (app.globalData?.customBarHeight || 88) + 8
			}
		},
		onShow() { this.load() },
		watch: {
			currentPeriod() { this.load() }
		},
		methods: {
			async load() {
				try {
					const res = await getStats(this.currentPeriod)
					this.stats = res.data || {}
				} catch (e) {}
			},
			formatDay(d) {
				if (!d) return ''
				const date = new Date(d)
				return (date.getMonth()+1) + '/' + date.getDate()
			},
			getBarWidth(weight) {
				const weights = (this.stats.weight_trend || []).map(w => w.weight)
				const min = Math.min(...weights) - 2
				const max = Math.max(...weights) + 2
				return ((weight - min) / (max - min)) * 100
			},
			getWorkoutBarHeight(count) {
				const max = Math.max(...(this.stats.workout_stats || []).map(s => s.count), 1)
				return Math.max((count / max) * 180, 8) + 'rpx'
			},
			getDietBarHeight(calories) {
				const max = Math.max(...(this.stats.diet_stats || []).map(s => s.calories), 1)
				return Math.max((calories / max) * 180, 8) + 'rpx'
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
	margin-bottom: $spacing-lg;

	.page-title {
		@include page-title;
	}
}

// --- Period Selector ---
.period-bar {
	display: flex;
	background: $color-fill;
	border-radius: $radius-md;
	padding: 4rpx;
	margin-bottom: $spacing-xl;

	.period-item {
		flex: 1;
		text-align: center;
		padding: $spacing-sm 0;
		font-size: $font-footnote;
		font-weight: 500;
		color: $color-label-tertiary;
		border-radius: calc(#{$radius-md} - 4rpx);
		transition: all 0.3s cubic-bezier(0.25, 0.46, 0.45, 0.94);

		&.active {
			background: $color-bg-elevated;
			color: $color-label;
			font-weight: 600;
			box-shadow: 0 1rpx 8rpx rgba(0, 0, 0, 0.08), 0 2rpx 16rpx rgba(0, 0, 0, 0.04);
		}
	}
}

// --- Stat Card ---
.stat-card {
	@include card;
	margin-bottom: $spacing-md;
	padding: $spacing-xl;

	.stat-header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: $spacing-lg;

		.stat-header-left {
			display: flex;
			align-items: center;
			gap: $spacing-sm;
		}

		.stat-dot {
			width: 16rpx;
			height: 16rpx;
			border-radius: 50%;
		}
		.dot-blue { background: $color-primary; }
		.dot-green { background: $color-green; }
		.dot-orange { background: $color-orange; }

		.stat-label {
			font-size: $font-headline;
			font-weight: 600;
			color: $color-label;
		}
		.stat-icon {
			font-size: 28rpx;
		}
	}
}

.no-data {
	padding: $spacing-2xl 0;
	text-align: center;

	.no-data-text {
		font-size: $font-subhead;
		color: $color-separator-opaque;
	}
}

// --- Bar Chart ---
.chart-area {
	padding-top: $spacing-md;
}

.bar-chart {
	display: flex;
	align-items: flex-end;
	justify-content: space-around;
	height: 260rpx;

	.bar-col {
		flex: 1;
		display: flex;
		flex-direction: column;
		align-items: center;

		.bar-val {
			font-size: $font-caption2;
			color: $color-label-tertiary;
			margin-bottom: $spacing-xs;
			font-weight: 600;
			font-variant-numeric: tabular-nums;
		}

		.bar-track {
			width: 36rpx;
			background: $color-fill;
			border-radius: 18rpx;
			overflow: hidden;
			display: flex;
			align-items: flex-end;
			min-height: 8rpx;
		}

		.bar-fill {
			width: 100%;
			border-radius: 18rpx;
			min-height: 8rpx;
			transition: height 0.5s cubic-bezier(0.25, 0.46, 0.45, 0.94);
		}
		.fill-blue {
			background: linear-gradient(180deg, $color-primary, #0062cc);
		}
		.fill-orange {
			background: linear-gradient(180deg, $color-orange, #e07800);
		}

		.bar-day {
			font-size: $font-caption2;
			color: $color-label-quaternary;
			margin-top: $spacing-sm;
			font-weight: 500;
		}
	}
}

// --- Trend List ---
.trend-area {
	padding-top: $spacing-xs;
}

.trend-row {
	display: flex;
	align-items: center;
	margin-bottom: $spacing-md;

	&:last-child { margin-bottom: 0; }

	.trend-day {
		width: 80rpx;
		font-size: $font-caption1;
		color: $color-label-quaternary;
		font-weight: 500;
		font-variant-numeric: tabular-nums;
	}

	.trend-bar-track {
		flex: 1;
		height: 24rpx;
		background: $color-fill;
		border-radius: 12rpx;
		margin: 0 $spacing-md;
		overflow: hidden;

		.trend-bar-fill {
			height: 100%;
			background: linear-gradient(90deg, $color-green, $color-mint);
			border-radius: 12rpx;
			min-width: 8rpx;
			transition: width 0.5s cubic-bezier(0.25, 0.46, 0.45, 0.94);
		}
	}

	.trend-val {
		width: 120rpx;
		text-align: right;
		font-size: $font-subhead;
		font-weight: 600;
		color: $color-label;
		font-variant-numeric: tabular-nums;

		.trend-unit {
			font-size: $font-caption1;
			color: $color-label-quaternary;
			font-weight: 400;
		}
	}
}
</style>
