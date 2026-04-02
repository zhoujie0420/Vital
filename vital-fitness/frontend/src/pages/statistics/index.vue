<template>
	<view class="container">
		<!-- 顶部导航栏 -->
		<u-navbar title="数据统计" fixed placeholder>
			<view slot="right">
				<u-icon name="calendar" size="20" @click="showDateRangePicker"></u-icon>
			</view>
		</u-navbar>

		<!-- 时间范围选择 -->
		<view class="date-range">
			<u-datetime-picker
				:show="showStartDatePicker"
				v-model="dateRange.start"
				mode="date"
				@confirm="confirmStartDate"
				@cancel="showStartDatePicker = false"
			></u-datetime-picker>
			<u-datetime-picker
				:show="showEndDatePicker"
				v-model="dateRange.end"
				mode="date"
				@confirm="confirmEndDate"
				@cancel="showEndDatePicker = false"
			></u-datetime-picker>

			<view class="date-range-selector" @click="selectDateRange">
				<text>{{ formatDateRange }}</text>
				<u-icon name="arrow-down" size="16" color="#999"></u-icon>
			</view>
		</view>

		<!-- 统计概览 -->
		<view class="stats-overview">
			<u-row gutter="20">
				<u-col span="6">
					<view class="overview-card">
						<text class="card-title">训练次数</text>
						<text class="card-value">{{ overviewStats.workoutCount }}</text>
						<text class="card-change" :class="workoutChangeClass">
							<u-icon :name="workoutChangeIcon" :color="workoutChangeColor" size="14"></u-icon>
							{{ Math.abs(overviewStats.workoutChange) }}%
						</text>
					</view>
				</u-col>
				<u-col span="6">
					<view class="overview-card">
						<text class="card-title">总消耗</text>
						<text class="card-value">{{ overviewStats.totalCalories }}</text>
						<text class="card-unit">千卡</text>
						<text class="card-change" :class="caloriesChangeClass">
							<u-icon :name="caloriesChangeIcon" :color="caloriesChangeColor" size="14"></u-icon>
							{{ Math.abs(overviewStats.caloriesChange) }}%
						</text>
					</view>
				</u-col>
			</u-row>

			<u-row gutter="20">
				<u-col span="6">
					<view class="overview-card">
						<text class="card-title">平均体重</text>
						<text class="card-value">{{ overviewStats.avgWeight }}</text>
						<text class="card-unit">kg</text>
						<text class="card-change" :class="weightChangeClass">
							<u-icon :name="weightChangeIcon" :color="weightChangeColor" size="14"></u-icon>
							{{ Math.abs(overviewStats.weightChange) }}kg
						</text>
					</view>
				</u-col>
				<u-col span="6">
					<view class="overview-card">
						<text class="card-title">平均心情</text>
						<view class="mood-display">
							<u-rate :value="overviewStats.avgMood" readonly size="14" active-color="#ff6b6b"></u-rate>
						</view>
						<text class="card-value small">{{ overviewStats.avgMood }}</text>
						<text class="card-unit">分</text>
					</view>
				</u-col>
			</u-row>
		</view>

		<!-- 图表区域 -->
		<view class="charts-section">
			<!-- 训练趋势图 -->
			<view class="chart-card">
				<view class="chart-header">
					<text class="chart-title">训练趋势</text>
					<u-icon name="reload" size="18" color="#999" @click="refreshWorkoutChart"></u-icon>
				</view>
				<view class="chart-container">
					<!-- 简化的训练趋势图 -->
					<view class="simple-chart">
						<view class="chart-grid">
							<view class="grid-lines">
								<view class="grid-line" v-for="i in 5" :key="i"></view>
							</view>
							<view class="chart-bars">
								<view
									v-for="(bar, index) in workoutTrend"
									:key="index"
									class="chart-bar"
									:style="{ height: bar.height }"
								>
									<text class="bar-value">{{ bar.value }}</text>
								</view>
							</view>
							<view class="chart-labels">
								<text v-for="(label, index) in trendLabels" :key="index" class="chart-label">{{ label }}</text>
							</view>
						</view>
					</view>
				</view>
			</view>

			<!-- 体重变化图 -->
			<view class="chart-card">
				<view class="chart-header">
					<text class="chart-title">体重变化</text>
				</view>
				<view class="chart-container">
					<!-- 简化的体重变化图 -->
					<view class="simple-line-chart">
						<view class="line-chart-grid">
							<view class="grid-lines">
								<view class="horizontal-line" v-for="i in 5" :key="i"></view>
							</view>
							<view class="line-path">
								<view class="line-point" v-for="(point, index) in weightPoints" :key="index"
									:style="{ left: point.left, bottom: point.bottom }"></view>
							</view>
						</view>
					</view>
				</view>
			</view>

			<!-- 饮食营养图 -->
			<view class="chart-card">
				<view class="chart-header">
					<text class="chart-title">营养摄入</text>
				</view>
				<view class="chart-container">
					<!-- 简化的营养摄入饼图 -->
					<view class="simple-pie-chart">
						<view class="pie-chart-container">
							<view class="pie-segment" v-for="(segment, index) in nutritionSegments" :key="index"
								:style="{ backgroundColor: segment.color, transform: `rotate(${segment.rotate}deg) skew(${segment.skew}deg)` }">
							</view>
						</view>
						<view class="pie-legend">
							<view class="legend-item" v-for="(segment, index) in nutritionSegments" :key="index">
								<view class="legend-color" :style="{ backgroundColor: segment.color }"></view>
								<text class="legend-label">{{ segment.label }} {{ segment.percent }}%</text>
							</view>
						</view>
					</view>
				</view>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				showStartDatePicker: false,
				showEndDatePicker: false,

				dateRange: {
					start: new Date(Date.now() - 30 * 24 * 60 * 60 * 1000).getTime(), // 30天前
					end: Date.now()
				},

				overviewStats: {
					workoutCount: 25,
					workoutChange: 12,
					totalCalories: 12500,
					caloriesChange: 8,
					avgWeight: 70.5,
					weightChange: -0.8,
					avgMood: 7.2
				},

				workoutTrend: [
					{ height: '60%', value: 3 },
					{ height: '80%', value: 4 },
					{ height: '40%', value: 2 },
					{ height: '100%', value: 5 },
					{ height: '70%', value: 3 },
					{ height: '90%', value: 4 },
					{ height: '50%', value: 2 }
				],

				trendLabels: ['周一', '周二', '周三', '周四', '周五', '周六', '周日'],

				weightPoints: [
					{ left: '10%', bottom: '30%' },
					{ left: '25%', bottom: '40%' },
					{ left: '40%', bottom: '35%' },
					{ left: '55%', bottom: '50%' },
					{ left: '70%', bottom: '45%' },
					{ left: '85%', bottom: '60%' }
				],

				nutritionSegments: [
					{ label: '蛋白质', percent: 25, color: '#3c9cff', rotate: 0, skew: 36 },
					{ label: '碳水化合物', percent: 50, color: '#ff9900', rotate: 36, skew: 36 },
					{ label: '脂肪', percent: 25, color: '#6bcf7f', rotate: 72, skew: 36 }
				]
			}
		},
		computed: {
			formatDateRange() {
				const startDate = new Date(this.dateRange.start)
				const endDate = new Date(this.dateRange.end)
				return `${startDate.getMonth() + 1}.${startDate.getDate()} - ${endDate.getMonth() + 1}.${endDate.getDate()}`
			},
			workoutChangeClass() {
				return this.overviewStats.workoutChange >= 0 ? 'change-up' : 'change-down'
			},
			workoutChangeIcon() {
				return this.overviewStats.workoutChange >= 0 ? 'arrow-up' : 'arrow-down'
			},
			workoutChangeColor() {
				return this.overviewStats.workoutChange >= 0 ? '#ff6b6b' : '#6bcf7f'
			},
			caloriesChangeClass() {
				return this.overviewStats.caloriesChange >= 0 ? 'change-up' : 'change-down'
			},
			caloriesChangeIcon() {
				return this.overviewStats.caloriesChange >= 0 ? 'arrow-up' : 'arrow-down'
			},
			caloriesChangeColor() {
				return this.overviewStats.caloriesChange >= 0 ? '#ff6b6b' : '#6bcf7f'
			},
			weightChangeClass() {
				return this.overviewStats.weightChange <= 0 ? 'change-down' : 'change-up'
			},
			weightChangeIcon() {
				return this.overviewStats.weightChange <= 0 ? 'arrow-down' : 'arrow-up'
			},
			weightChangeColor() {
				return this.overviewStats.weightChange <= 0 ? '#6bcf7f' : '#ff6b6b'
			}
		},
		methods: {
			showDateRangePicker() {
				// 显示日期范围选择器
				uni.showActionSheet({
					itemList: ['近7天', '近30天', '近90天', '自定义'],
					success: (res) => {
						const ranges = [7, 30, 90]
						if (res.tapIndex < 3) {
							const days = ranges[res.tapIndex]
							this.dateRange.end = Date.now()
							this.dateRange.start = new Date(Date.now() - days * 24 * 60 * 60 * 1000).getTime()
						} else {
							// 自定义日期范围
							this.showStartDatePicker = true
						}
					}
				})
			},

			confirmStartDate(e) {
				this.dateRange.start = e.value
				this.showStartDatePicker = false
				this.showEndDatePicker = true
			},

			confirmEndDate(e) {
				this.dateRange.end = e.value
				this.showEndDatePicker = false
			},

			selectDateRange() {
				this.showDateRangePicker()
			},

			refreshWorkoutChart() {
				// 刷新训练趋势图
				uni.showToast({
					title: '图表已刷新',
					icon: 'success'
				})
			}
		}
	}
</script>

<style lang="scss" scoped>
.container {
	padding: 20rpx;
}

.date-range {
	background: white;
	border-radius: 20rpx;
	padding: 20rpx 30rpx;
	margin-bottom: 20rpx;

	.date-range-selector {
		display: flex;
		align-items: center;
		justify-content: center;

		text {
			font-size: 28rpx;
			color: #333;
			margin-right: 10rpx;
		}
	}
}

.stats-overview {
	margin-bottom: 20rpx;

	.overview-card {
		background: white;
		border-radius: 20rpx;
		padding: 30rpx;
		text-align: center;
		height: 180rpx;
		display: flex;
		flex-direction: column;
		justify-content: center;

		.card-title {
			font-size: 24rpx;
			color: #999;
			margin-bottom: 10rpx;
		}

		.card-value {
			font-size: 36rpx;
			font-weight: bold;
			color: #333;
			margin-bottom: 5rpx;

			&.small {
				font-size: 28rpx;
			}
		}

		.card-unit {
			font-size: 24rpx;
			color: #999;
			margin-bottom: 5rpx;
		}

		.mood-display {
			margin-bottom: 10rpx;
			display: flex;
			justify-content: center;
		}

		.card-change {
			font-size: 20rpx;
			display: flex;
			align-items: center;
			justify-content: center;

			&.change-up {
				color: #ff6b6b;
			}

			&.change-down {
				color: #6bcf7f;
			}
		}
	}
}

.charts-section {
	.chart-card {
		background: white;
		border-radius: 20rpx;
		padding: 30rpx;
		margin-bottom: 20rpx;

		.chart-header {
			display: flex;
			justify-content: space-between;
			align-items: center;
			margin-bottom: 30rpx;

			.chart-title {
				font-size: 32rpx;
				font-weight: bold;
				color: #333;
			}
		}

		.chart-container {
			.simple-chart {
				.chart-grid {
					position: relative;
					height: 300rpx;

					.grid-lines {
						position: absolute;
						top: 0;
						left: 0;
						right: 0;
						bottom: 0;
						display: flex;
						flex-direction: column;
						justify-content: space-between;

						.grid-line {
							height: 1rpx;
							background: #eee;
						}
					}

					.chart-bars {
						position: absolute;
						top: 0;
						left: 0;
						right: 0;
						bottom: 60rpx;
						display: flex;
						align-items: end;
						justify-content: space-around;
						padding: 0 30rpx;

						.chart-bar {
							width: 40rpx;
							background: linear-gradient(to top, #3c9cff, #6bb9ff);
							border-radius: 10rpx 10rpx 0 0;
							position: relative;
							display: flex;
							align-items: start;
							justify-content: center;

							.bar-value {
								font-size: 20rpx;
								color: #333;
								margin-top: 10rpx;
							}
						}
					}

					.chart-labels {
						position: absolute;
						left: 0;
						right: 0;
						bottom: 0;
						height: 60rpx;
						display: flex;
						justify-content: space-around;
						align-items: center;

						.chart-label {
							font-size: 20rpx;
							color: #999;
						}
					}
				}
			}

			.simple-line-chart {
				.line-chart-grid {
					position: relative;
					height: 300rpx;

					.grid-lines {
						position: absolute;
						top: 0;
						left: 0;
						right: 0;
						bottom: 0;
						display: flex;
						flex-direction: column;
						justify-content: space-between;

						.horizontal-line {
							height: 1rpx;
							background: #eee;
						}
					}

					.line-path {
						position: absolute;
						top: 0;
						left: 0;
						right: 0;
						bottom: 0;

						.line-point {
							position: absolute;
							width: 12rpx;
							height: 12rpx;
							background: #3c9cff;
							border-radius: 50%;
							transform: translate(-50%, 50%);
						}

						// 简化的连线效果
						&::after {
							content: '';
							position: absolute;
							top: 0;
							left: 10%;
							width: 75%;
							height: 2rpx;
							background: linear-gradient(to right, #3c9cff, #6bb9ff, #99d9ff);
							bottom: 30%;
						}
					}
				}
			}

			.simple-pie-chart {
				display: flex;
				align-items: center;

				.pie-chart-container {
					width: 200rpx;
					height: 200rpx;
					position: relative;
					margin-right: 30rpx;

					.pie-segment {
						position: absolute;
						width: 200rpx;
						height: 200rpx;
						border-radius: 50%;
						clip-path: polygon(50% 50%, 50% 0%, 100% 0%, 100% 100%, 50% 100%);
						transform-origin: center;
					}
				}

				.pie-legend {
					flex: 1;

					.legend-item {
						display: flex;
						align-items: center;
						margin-bottom: 20rpx;

						.legend-color {
							width: 20rpx;
							height: 20rpx;
							border-radius: 4rpx;
							margin-right: 15rpx;
						}

						.legend-label {
							font-size: 24rpx;
							color: #666;
						}
					}
				}
			}
		}
	}
}
</style>