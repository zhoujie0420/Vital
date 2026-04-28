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
					<view class="stat-dot dot-red"></view>
					<text class="stat-label">训练次数</text>
				</view>
				<text class="stat-summary" v-if="workoutTotal > 0">共 {{ workoutTotal }} 次</text>
			</view>
			<view v-if="stats.workout_stats && stats.workout_stats.length > 0" class="chart-area">
				<view class="bar-chart">
					<view v-for="(s, i) in stats.workout_stats" :key="i" class="bar-col">
						<text class="bar-val" v-if="s.count > 0">{{ s.count }}</text>
						<view class="bar-track">
							<view class="bar-fill fill-red" :style="{ height: getWorkoutBarHeight(s.count) }"></view>
						</view>
						<text class="bar-day">{{ formatDay(s.day) }}</text>
					</view>
				</view>
			</view>
			<view v-else class="no-data">
				<view class="no-data-icon">
					<text class="no-data-letter">W</text>
				</view>
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
				<text class="stat-summary" v-if="stats.weight_trend && stats.weight_trend.length > 0">
					{{ stats.weight_trend[0].weight }} kg
				</text>
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
				<view class="no-data-icon">
					<text class="no-data-letter">B</text>
				</view>
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
				<text class="stat-summary" v-if="dietTotal > 0">日均 {{ dietAvg }} kcal</text>
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
				<view class="no-data-icon">
					<text class="no-data-letter">D</text>
				</view>
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
			},
			workoutTotal() {
				return (this.stats.workout_stats || []).reduce((s, i) => s + i.count, 0)
			},
			dietTotal() {
				return (this.stats.diet_stats || []).reduce((s, i) => s + i.calories, 0)
			},
			dietAvg() {
				const list = (this.stats.diet_stats || []).filter(i => i.calories > 0)
				if (list.length === 0) return 0
				return Math.round(list.reduce((s, i) => s + i.calories, 0) / list.length)
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
	.page-title { @include page-title; }
}

// --- Period Selector ---
.period-bar {
	@include segment-control;
	margin-bottom: $spacing-xl;
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
			width: 14rpx;
			height: 14rpx;
			border-radius: 50%;
		}
		.dot-red { background: $color-red; }
		.dot-green { background: $color-green; }
		.dot-orange { background: $color-accent; }

		.stat-label {
			font-size: $font-headline;
			font-weight: 600;
			color: $color-label;
			letter-spacing: -0.3rpx;
		}
		.stat-summary {
			font-size: $font-caption1;
			color: $color-label-quaternary;
			font-weight: 500;
			font-variant-numeric: tabular-nums;
		}
	}
}

.no-data {
	padding: $spacing-2xl 0;
	text-align: center;

	.no-data-icon {
		width: 72rpx;
		height: 72rpx;
		border-radius: 50%;
		background: $color-fill;
		display: flex;
		align-items: center;
		justify-content: center;
		margin: 0 auto $spacing-sm;
	}
	.no-data-letter {
		font-size: $font-headline;
		font-weight: 700;
		color: $color-label-quaternary;
	}
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
			width: 32rpx;
			background: $color-fill;
			border-radius: 16rpx;
			overflow: hidden;
			display: flex;
			align-items: flex-end;
			min-height: 8rpx;
		}

		.bar-fill {
			width: 100%;
			border-radius: 16rpx;
			min-height: 8rpx;
			transition: height 0.5s cubic-bezier(0.16, 1, 0.3, 1);
		}
		.fill-red {
			background: linear-gradient(180deg, $color-red, #DC2626);
		}
		.fill-orange {
			background: linear-gradient(180deg, $color-accent, #EA580C);
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
		height: 20rpx;
		background: $color-fill;
		border-radius: 10rpx;
		margin: 0 $spacing-md;
		overflow: hidden;

		.trend-bar-fill {
			height: 100%;
			background: linear-gradient(90deg, $color-green, $color-mint);
			border-radius: 10rpx;
			min-width: 8rpx;
			transition: width 0.5s cubic-bezier(0.16, 1, 0.3, 1);
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
