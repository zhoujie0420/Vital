<template>
	<view class="container">
		<u-navbar title="数据统计" fixed placeholder />

		<!-- 训练统计 -->
		<view class="stat-card">
			<text class="card-title">近7天训练</text>
			<view class="bar-chart">
				<view v-for="(s, i) in stats.workout_stats" :key="i" class="bar-item">
					<view class="bar" :style="{ height: (s.count * 40 + 10) + 'rpx' }"></view>
					<text class="bar-label">{{ formatDay(s.day) }}</text>
				</view>
			</view>
			<text v-if="!stats.workout_stats || stats.workout_stats.length === 0" class="no-data">暂无训练数据</text>
		</view>

		<!-- 体重趋势 -->
		<view class="stat-card">
			<text class="card-title">体重趋势</text>
			<view v-if="stats.weight_trend && stats.weight_trend.length > 0" class="weight-list">
				<view v-for="(w, i) in stats.weight_trend" :key="i" class="weight-row">
					<text class="weight-day">{{ formatDay(w.day) }}</text>
					<text class="weight-val">{{ w.weight }}kg</text>
				</view>
			</view>
			<text v-else class="no-data">暂无体重数据</text>
		</view>

		<!-- 饮食热量 -->
		<view class="stat-card">
			<text class="card-title">近7天热量摄入</text>
			<view class="bar-chart">
				<view v-for="(s, i) in stats.diet_stats" :key="i" class="bar-item">
					<view class="bar bar-orange" :style="{ height: Math.min(s.calories / 10, 200) + 'rpx' }"></view>
					<text class="bar-label">{{ formatDay(s.day) }}</text>
				</view>
			</view>
			<text v-if="!stats.diet_stats || stats.diet_stats.length === 0" class="no-data">暂无饮食数据</text>
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
			}
		}
	}
</script>

<style lang="scss" scoped>
.container { padding: 20rpx; }
.stat-card {
	background: white; border-radius: 20rpx; padding: 30rpx; margin-bottom: 20rpx;
	.card-title { font-size: 32rpx; font-weight: bold; color: #333; display: block; margin-bottom: 20rpx; }
	.no-data { font-size: 26rpx; color: #999; text-align: center; display: block; padding: 30rpx 0; }
	.bar-chart {
		display: flex; align-items: flex-end; justify-content: space-around; height: 250rpx; padding-top: 20rpx;
		.bar-item { text-align: center; flex: 1;
			.bar { width: 40rpx; margin: 0 auto; background: #3c9cff; border-radius: 8rpx 8rpx 0 0; min-height: 10rpx; }
			.bar-orange { background: #ff9900; }
			.bar-label { font-size: 20rpx; color: #999; margin-top: 8rpx; display: block; }
		}
	}
	.weight-list {
		.weight-row { display: flex; justify-content: space-between; padding: 15rpx 0; border-bottom: 1rpx solid #f5f5f5;
			.weight-day { font-size: 26rpx; color: #999; }
			.weight-val { font-size: 28rpx; font-weight: bold; color: #333; }
		}
	}
}
</style>
