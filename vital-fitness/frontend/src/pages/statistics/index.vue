<template>
	<view class="page">
		<view class="page-header">
			<text class="page-title">数据统计</text>
		</view>

		<!-- 训练统计 -->
		<view class="card">
			<text class="card-label">近7天训练</text>
			<view v-if="stats.workout_stats && stats.workout_stats.length > 0" class="bar-chart">
				<view v-for="(s, i) in stats.workout_stats" :key="i" class="bar-col">
					<view class="bar-fill" :style="{ height: (s.count * 40 + 10) + 'rpx' }"></view>
					<text class="bar-val" v-if="s.count > 0">{{ s.count }}</text>
					<text class="bar-day">{{ formatDay(s.day) }}</text>
				</view>
			</view>
			<text v-else class="no-data">暂无训练数据</text>
		</view>

		<!-- 体重趋势 -->
		<view class="card">
			<text class="card-label">体重趋势</text>
			<view v-if="stats.weight_trend && stats.weight_trend.length > 0" class="trend-list">
				<view v-for="(w, i) in stats.weight_trend" :key="i" class="trend-row">
					<text class="trend-day">{{ formatDay(w.day) }}</text>
					<view class="trend-bar-wrap">
						<view class="trend-bar" :style="{ width: getBarWidth(w.weight) + '%' }"></view>
					</view>
					<text class="trend-val">{{ w.weight }}kg</text>
				</view>
			</view>
			<text v-else class="no-data">暂无体重数据</text>
		</view>

		<!-- 饮食热量 -->
		<view class="card">
			<text class="card-label">近7天热量摄入</text>
			<view v-if="stats.diet_stats && stats.diet_stats.length > 0" class="bar-chart">
				<view v-for="(s, i) in stats.diet_stats" :key="i" class="bar-col">
					<view class="bar-fill bar-orange" :style="{ height: Math.min(s.calories / 10, 200) + 'rpx' }"></view>
					<text class="bar-val" v-if="s.calories > 0">{{ s.calories }}</text>
					<text class="bar-day">{{ formatDay(s.day) }}</text>
				</view>
			</view>
			<text v-else class="no-data">暂无饮食数据</text>
		</view>
	</view>
</template>

<script>
	import { getStats } from '../../api/stats'

	export default {
		data() { return { stats: {} } },
		onShow() { this.load() },
		methods: {
			async load() {
				try {
					const res = await getStats()
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
			}
		}
	}
</script>

<style lang="scss" scoped>
.page {
	padding: 0 32rpx; padding-top: 120rpx; padding-bottom: 40rpx;
	min-height: 100vh; background: #f2f2f7;
}

.page-header {
	margin-bottom: 28rpx;
	.page-title {
		font-size: 52rpx; font-weight: 700; color: #1c1c1e; letter-spacing: -1rpx;
	}
}

.card {
	background: #fff; border-radius: 20rpx; padding: 28rpx 32rpx; margin-bottom: 16rpx;
	box-shadow: 0 2rpx 16rpx rgba(0, 0, 0, 0.04);
	.card-label {
		display: block; font-size: 26rpx; font-weight: 600; color: #8e8e93;
		text-transform: uppercase; letter-spacing: 1rpx; margin-bottom: 24rpx;
	}
}

.no-data {
	display: block; text-align: center; font-size: 28rpx; color: #c7c7cc;
	padding: 40rpx 0;
}

.bar-chart {
	display: flex; align-items: flex-end; justify-content: space-around;
	height: 280rpx; padding-top: 40rpx;

	.bar-col {
		flex: 1; display: flex; flex-direction: column; align-items: center;

		.bar-fill {
			width: 36rpx; background: #007aff; border-radius: 8rpx 8rpx 0 0;
			min-height: 8rpx;
		}
		.bar-orange { background: #ff9500; }

		.bar-val {
			font-size: 20rpx; color: #8e8e93; margin-bottom: 6rpx; font-weight: 600;
		}
		.bar-day {
			font-size: 22rpx; color: #8e8e93; margin-top: 12rpx;
		}
	}
}

.trend-list {
	.trend-row {
		display: flex; align-items: center; margin-bottom: 16rpx;

		.trend-day { width: 80rpx; font-size: 24rpx; color: #8e8e93; }
		.trend-bar-wrap {
			flex: 1; height: 24rpx; background: #f2f2f7; border-radius: 12rpx;
			margin: 0 16rpx; overflow: hidden;

			.trend-bar {
				height: 100%; background: linear-gradient(90deg, #34c759, #30d158);
				border-radius: 12rpx; min-width: 8rpx;
			}
		}
		.trend-val { width: 120rpx; text-align: right; font-size: 28rpx; font-weight: 600; color: #1c1c1e; }
	}
}
</style>
