<template>
	<view class="container">
		<!-- 顶部导航栏 -->
		<u-navbar title="体重记录" fixed placeholder>
			<view slot="right">
				<u-icon name="plus" size="20" @click="addWeight"></u-icon>
			</view>
		</u-navbar>

		<!-- 体重统计卡片 -->
		<view class="weight-stats">
			<view class="stats-card">
				<text class="card-title">当前体重</text>
				<view class="weight-display">
					<text class="weight-value">{{ currentWeight }}kg</text>
				</view>
				<text class="weight-change" :class="weightChangeClass">
					<u-icon :name="weightChangeIcon" :color="weightChangeColor" size="16"></u-icon>
					{{ Math.abs(weightChange) }}kg {{ weightChangeText }}
				</text>
			</view>

			<view class="stats-card">
				<text class="card-title">BMI指数</text>
				<view class="bmi-display">
					<text class="bmi-value">{{ currentBMI }}</text>
				</view>
				<text class="bmi-status" :class="bmiStatusClass">{{ bmiStatus }}</text>
			</view>
		</view>

		<!-- 体重趋势图 -->
		<view class="weight-chart">
			<view class="chart-header">
				<text class="chart-title">体重变化趋势</text>
				<text class="chart-subtitle">近30天</text>
			</view>

			<!-- 简化的体重趋势图 -->
			<view class="weight-trend">
				<view class="trend-line">
					<view v-for="(point, index) in weightTrend" :key="index" class="trend-point"
						:style="{ height: point.height, backgroundColor: point.color }">
					</view>
				</view>

				<view class="trend-dates">
					<text v-for="(date, index) in trendDates" :key="index" class="date-item">{{ date }}</text>
				</view>
			</view>
		</view>

		<!-- 体重记录列表 -->
		<view class="weight-list">
			<view class="list-header">
				<text class="header-title">历史记录</text>
			</view>

			<view v-for="(weight, index) in weightList" :key="index" class="weight-item">
				<view class="weight-header">
					<text class="weight-date">{{ formatDate(weight.record_date) }}</text>
					<text class="weight-time">{{ formatTime(weight.created_at) }}</text>
				</view>

				<view class="weight-content">
					<view class="weight-info">
						<text class="weight-value">{{ weight.weight }}kg</text>
						<text v-if="weight.height" class="height-value">{{ weight.height }}cm</text>
					</view>

					<view v-if="weight.bmi" class="bmi-info">
						<text class="bmi-label">BMI:</text>
						<text class="bmi-value">{{ weight.bmi }}</text>
						<text class="bmi-status-small" :class="weight.bmi < 18.5 ? 'status-underweight' : weight.bmi < 24 ? 'status-normal' : weight.bmi < 28 ? 'status-overweight' : 'status-obese'">{{ getBMIStatus(weight.bmi) }}</text>
					</view>
				</view>
			</view>

			<!-- 空状态 -->
			<view v-if="weightList.length === 0" class="empty-state">
				<u-empty text="暂无体重记录" mode="weight"></u-empty>
				<u-button type="primary" icon="plus" @click="addWeight" style="margin-top: 30rpx;">记录第一次体重</u-button>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				currentWeight: 70.5,
				previousWeight: 71.2,
				currentBMI: 22.1,
				weightList: [
					{
						id: 1,
						record_date: '2026-03-12',
						created_at: '2026-03-12 08:00:00',
						weight: 70.5,
						height: 178,
						bmi: 22.1
					},
					{
						id: 2,
						record_date: '2026-03-05',
						created_at: '2026-03-05 08:00:00',
						weight: 71.2,
						height: 178,
						bmi: 22.3
					},
					{
						id: 3,
						record_date: '2026-02-28',
						created_at: '2026-02-28 08:00:00',
						weight: 71.8,
						height: 178,
						bmi: 22.5
					}
				],
				weightTrend: [
					{ height: '40rpx', color: '#3c9cff' },
					{ height: '60rpx', color: '#5ca9ff' },
					{ height: '80rpx', color: '#7db3ff' },
					{ height: '100rpx', color: '#9ebdff' },
					{ height: '120rpx', color: '#bfc7ff' },
					{ height: '140rpx', color: '#dfd1ff' },
					{ height: '160rpx', color: '#ffcfcc' }
				],
				trendDates: ['2/1', '2/10', '2/20', '3/1', '3/5', '3/10', '3/12']
			}
		},
		computed: {
			weightChange() {
				return this.currentWeight - this.previousWeight
			},
			weightChangeClass() {
				return this.weightChange < 0 ? 'change-down' : this.weightChange > 0 ? 'change-up' : 'change-none'
			},
			weightChangeIcon() {
				return this.weightChange < 0 ? 'arrow-down' : this.weightChange > 0 ? 'arrow-up' : 'minus'
			},
			weightChangeColor() {
				return this.weightChange < 0 ? '#6bcf7f' : this.weightChange > 0 ? '#ff6b6b' : '#999'
			},
			weightChangeText() {
				return this.weightChange < 0 ? '下降' : this.weightChange > 0 ? '上升' : '无变化'
			},
			bmiStatus() {
				return this.getBMIStatus(this.currentBMI)
			},
			bmiStatusClass() {
				return this.getBMIStatusClass(this.currentBMI)
			}
		},
		methods: {
			addWeight() {
				uni.navigateTo({
					url: '/pages/weight/add'
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
			},

			getBMIStatus(bmi) {
				if (bmi < 18.5) return '偏瘦'
				if (bmi < 24) return '正常'
				if (bmi < 28) return '超重'
				return '肥胖'
			},

			getBMIStatusClass(bmi) {
				if (bmi < 18.5) return 'status-underweight'
				if (bmi < 24) return 'status-normal'
				if (bmi < 28) return 'status-overweight'
				return 'status-obese'
			}
		}
	}
</script>

<style lang="scss" scoped>
.container {
	padding: 20rpx;
}

.weight-stats {
	display: flex;
	gap: 20rpx;
	margin-bottom: 20rpx;

	.stats-card {
		flex: 1;
		background: white;
		border-radius: 20rpx;
		padding: 30rpx;
		text-align: center;

		.card-title {
			font-size: 24rpx;
			color: #999;
			display: block;
			margin-bottom: 10rpx;
		}

		.weight-display,
		.bmi-display {
			margin: 20rpx 0;

			.weight-value {
				font-size: 48rpx;
				font-weight: bold;
				color: #333;
			}

			.bmi-value {
				font-size: 48rpx;
				font-weight: bold;
				color: #3c9cff;
			}
		}

		.weight-change {
			font-size: 24rpx;
			display: flex;
			align-items: center;
			justify-content: center;

			&.change-down {
				color: #6bcf7f;
			}

			&.change-up {
				color: #ff6b6b;
			}

			&.change-none {
				color: #999;
			}
		}

		.bmi-status {
			font-size: 24rpx;
			padding: 5rpx 15rpx;
			border-radius: 10rpx;
			display: inline-block;

			&.status-underweight {
				background: #ffeaa7;
				color: #d35400;
			}

			&.status-normal {
				background: #55a630;
				color: white;
			}

			&.status-overweight {
				background: #ff9f1c;
				color: white;
			}

			&.status-obese {
				background: #e63946;
				color: white;
			}
		}
	}
}

.weight-chart {
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
			display: block;
		}

		.chart-subtitle {
			font-size: 24rpx;
			color: #999;
		}
	}

	.weight-trend {
		.trend-line {
			display: flex;
			align-items: end;
			height: 200rpx;
			gap: 10rpx;
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

.weight-list {
	background: white;
	border-radius: 20rpx;
	overflow: hidden;

	.list-header {
		padding: 30rpx;
		border-bottom: 1rpx solid #eee;

		.header-title {
			font-size: 32rpx;
			font-weight: bold;
			color: #333;
		}
	}

	.weight-item {
		padding: 30rpx;
		border-bottom: 1rpx solid #eee;

		&:last-child {
			border-bottom: none;
		}

		.weight-header {
			display: flex;
			justify-content: space-between;
			align-items: center;
			margin-bottom: 20rpx;

			.weight-date {
				font-size: 28rpx;
				font-weight: bold;
				color: #333;
			}

			.weight-time {
				font-size: 24rpx;
				color: #999;
			}
		}

		.weight-content {
			.weight-info {
				display: flex;
				align-items: center;
				margin-bottom: 10rpx;

				.weight-value {
					font-size: 36rpx;
					font-weight: bold;
					color: #333;
					margin-right: 20rpx;
				}

				.height-value {
					font-size: 24rpx;
					color: #999;
				}
			}

			.bmi-info {
				display: flex;
				align-items: center;

				.bmi-label {
					font-size: 24rpx;
					color: #666;
					margin-right: 10rpx;
				}

				.bmi-value {
					font-size: 24rpx;
					color: #333;
					margin-right: 10rpx;
				}

				.bmi-status-small {
					font-size: 20rpx;
					padding: 2rpx 10rpx;
					border-radius: 5rpx;

					&.status-underweight {
						background: #ffeaa7;
						color: #d35400;
					}

					&.status-normal {
						background: #55a630;
						color: white;
					}

					&.status-overweight {
						background: #ff9f1c;
						color: white;
					}

					&.status-obese {
						background: #e63946;
						color: white;
					}
				}
			}
		}
	}

	.empty-state {
		text-align: center;
		padding: 100rpx 0;
	}
}
</style>