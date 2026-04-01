<template>
	<view class="container">
		<!-- 顶部导航栏 -->
		<u-navbar title="心情记录" fixed placeholder>
			<view slot="right">
				<u-icon name="plus" size="20" @click="addMood"></u-icon>
			</view>
		</u-navbar>

		<!-- 心情统计图表 -->
		<view class="mood-chart">
			<view class="chart-header">
				<text class="chart-title">近7天心情趋势</text>
			</view>

			<!-- 简化的心情趋势图（实际项目中可以使用图表库） -->
			<view class="mood-trend">
				<view class="trend-line">
					<view v-for="(point, index) in moodTrend" :key="index" class="trend-point"
						:style="{ height: point.height, backgroundColor: point.color }">
					</view>
				</view>

				<view class="trend-dates">
					<text v-for="(date, index) in trendDates" :key="index" class="date-item">{{ date }}</text>
				</view>
			</view>
		</view>

		<!-- 心情记录列表 -->
		<view class="mood-list">
			<view v-for="(mood, index) in moodList" :key="index" class="mood-item">
				<view class="mood-header">
					<text class="mood-date">{{ formatDate(mood.record_date) }}</text>
					<text class="mood-time">{{ formatTime(mood.created_at) }}</text>
				</view>

				<view class="mood-content">
					<view class="mood-score">
						<text class="score-label">心情评分</text>
						<view class="score-display">
							<text class="score-value">{{ mood.mood_score }}</text>
							<text class="score-max">/10</text>
						</view>
					</view>

					<u-rate :value="mood.mood_score" readonly size="16" active-color="#ff6b6b"></u-rate>

					<view v-if="mood.mood_tags" class="mood-tags">
						<u-tag v-for="tag in mood.mood_tags.split(',')" :key="tag" :text="tag" type="warning" size="mini"
							style="margin-right: 10rpx;" />
					</view>

					<view v-if="mood.description" class="mood-description">
						<text>{{ mood.description }}</text>
					</view>
				</view>
			</view>

			<!-- 空状态 -->
			<view v-if="moodList.length === 0" class="empty-state">
				<u-empty text="暂无心情记录" mode="favor"></u-empty>
				<u-button type="primary" icon="plus" @click="addMood" style="margin-top: 30rpx;">记录今天的心情</u-button>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				moodList: [
					{
						id: 1,
						record_date: '2026-03-12',
						created_at: '2026-03-12 09:00:00',
						mood_score: 8,
						mood_tags: '兴奋,满足',
						description: '今天训练很有成就感，感觉很棒！'
					},
					{
						id: 2,
						record_date: '2026-03-11',
						created_at: '2026-03-11 08:30:00',
						mood_score: 7,
						mood_tags: '平静,专注',
						description: '冥想之后感觉内心很平静'
					},
					{
						id: 3,
						record_date: '2026-03-10',
						created_at: '2026-03-10 19:00:00',
						mood_score: 6,
						mood_tags: '疲惫',
						description: '工作了一天，有点累但完成了训练'
					}
				],
				moodTrend: [
					{ height: '40rpx', color: '#ff6b6b' },
					{ height: '60rpx', color: '#ff8e8e' },
					{ height: '80rpx', color: '#ffb1b1' },
					{ height: '100rpx', color: '#ffd4d4' },
					{ height: '120rpx', color: '#ff8e8e' },
					{ height: '140rpx', color: '#ff6b6b' },
					{ height: '160rpx', color: '#ff4848' }
				],
				trendDates: ['3/7', '3/8', '3/9', '3/10', '3/11', '3/12', '3/13']
			}
		},
		methods: {
			addMood() {
				uni.navigateTo({
					url: '/pages/mood/add'
				})
			},

			formatDate(dateString) {
				const date = new Date(dateString)
				const today = new Date()

				if (date.toDateString() === today.toDateString()) {
					return '今天'
				}

				const yesterday = new Date(today)
				yesterday.setDate(yesterday.getDate() - 1)

				if (date.toDateString() === yesterday.toDateString()) {
					return '昨天'
				}

				return `${date.getMonth() + 1}月${date.getDate()}日`
			},

			formatTime(dateTimeString) {
				const date = new Date(dateTimeString)
				return `${date.getHours().toString().padStart(2, '0')}:${date.getMinutes().toString().padStart(2, '0')}`
			}
		}
	}
</script>

<style lang="scss" scoped>
.container {
	padding: 20rpx;
}

.mood-chart {
	background: white;
	border-radius: 20rpx;
	padding: 30rpx;
	margin-bottom: 20rpx;

	.chart-header {
		margin-bottom: 30rpx;

		.chart-title {
			font-size: 32rpx;
			font-weight: bold;
			color: #333;
		}
	}

	.mood-trend {
		.trend-line {
			display: flex;
			align-items: end;
			height: 200rpx;
			gap: 20rpx;
			margin-bottom: 20rpx;

			.trend-point {
				flex: 1;
				border-radius: 10rpx 10rpx 0 0;
				min-height: 10rpx;
			}
		}

		.trend-dates {
			display: flex;
			justify-content: space-between;

			.date-item {
				font-size: 20rpx;
				color: #999;
			}
		}
	}
}

.mood-list {
	.mood-item {
		background: white;
		border-radius: 20rpx;
		padding: 30rpx;
		margin-bottom: 20rpx;

		.mood-header {
			display: flex;
			justify-content: space-between;
			align-items: center;
			margin-bottom: 20rpx;

			.mood-date {
				font-size: 32rpx;
				font-weight: bold;
				color: #333;
			}

			.mood-time {
				font-size: 24rpx;
				color: #999;
			}
		}

		.mood-content {
			.mood-score {
				display: flex;
				justify-content: space-between;
				align-items: center;
				margin-bottom: 20rpx;

				.score-label {
					font-size: 28rpx;
					color: #333;
				}

				.score-display {
					.score-value {
						font-size: 48rpx;
						font-weight: bold;
						color: #ff6b6b;
					}

					.score-max {
						font-size: 24rpx;
						color: #999;
					}
				}
			}

			.mood-tags {
				margin: 20rpx 0;
			}

			.mood-description {
				margin-top: 20rpx;
				padding: 20rpx;
				background: #f5f5f5;
				border-radius: 10rpx;

				text {
					font-size: 26rpx;
					color: #666;
					line-height: 1.5;
				}
			}
		}
	}

	.empty-state {
		text-align: center;
		padding: 100rpx 0;
		background: white;
		border-radius: 20rpx;
	}
}
</style>