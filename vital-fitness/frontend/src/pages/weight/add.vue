<template>
	<view class="container">
		<!-- 顶部导航栏 -->
		<u-navbar title="记录体重" fixed placeholder>
			<view slot="left">
				<u-icon name="arrow-left" size="20" @click="goBack"></u-icon>
			</view>
			<view slot="right">
				<text class="nav-btn" @click="saveWeight">保存</text>
			</view>
		</u-navbar>

		<view class="form-container">
			<!-- 记录日期 -->
			<view class="form-item">
				<text class="label">记录日期</text>
				<u-datetime-picker
					:show="showDatePicker"
					v-model="weightForm.record_date"
					mode="date"
					@confirm="confirmDate"
					@cancel="showDatePicker = false"
				></u-datetime-picker>
				<u-cell
					title="记录日期"
					:value="formatDate(weightForm.record_date)"
					isLink
					@click="showDatePicker = true"
				></u-cell>
			</view>

			<!-- 体重输入 -->
			<view class="form-item">
				<text class="label">体重 (kg)</text>
				<view class="weight-input">
					<u-number-keyboard
						ref="keyboard"
						mode="number"
						:dotDisabled="false"
						:random="false"
						@change="keyboardChange"
						@backspace="keyboardBackspace"
					></u-number-keyboard>

					<view class="weight-display">
						<text class="weight-value">{{ weightDisplay }}</text>
						<text class="unit">kg</text>
					</view>
				</view>
			</view>

			<!-- 身高输入（可选） -->
			<view class="form-item">
				<text class="label">身高 (cm) <text class="optional">（可选）</text></text>
				<u-cell
					title="输入身高"
					:value="heightDisplay"
					isLink
					@click="showHeightInput = true"
				></u-cell>

				<!-- 身高输入弹窗 -->
				<u-popup :show="showHeightInput" mode="center" round closeable @close="showHeightInput = false">
					<view class="height-input-popup">
						<view class="popup-header">
							<text class="popup-title">输入身高</text>
						</view>

						<u-number-box
							v-model="tempHeight"
							:min="50"
							:max="250"
							:step="1"
							integer
							label="cm"
						></u-number-box>

						<view class="popup-buttons">
							<u-button type="info" @click="showHeightInput = false" style="margin-right: 20rpx;">取消</u-button>
							<u-button type="primary" @click="confirmHeight">确定</u-button>
						</view>
					</view>
				</u-popup>
			</view>

			<!-- BMI计算结果 -->
			<view class="form-item" v-if="calculatedBMI > 0">
				<text class="label">BMI计算结果</text>
				<view class="bmi-result">
					<text class="bmi-value">{{ calculatedBMI }}</text>
					<text class="bmi-status" :class="bmiStatusClass">{{ bmiStatus }}</text>
				</view>

				<view class="bmi-reference">
					<text class="reference-title">参考标准：</text>
					<view class="reference-items">
						<text class="reference-item"><text class="underweight">偏瘦</text> 小于 18.5</text>
						<text class="reference-item"><text class="normal">正常</text> 18.5-23.9</text>
						<text class="reference-item"><text class="overweight">超重</text> 24-27.9</text>
						<text class="reference-item"><text class="obese">肥胖</text> 大于等于 28</text>
					</view>
				</view>
			</view>

			<!-- 体重变化趋势 -->
			<view class="form-item" v-if="previousWeight > 0">
				<text class="label">体重变化</text>
				<view class="weight-change">
					<text>相比上次记录</text>
					<text class="change-value" :class="changeClass">
						<u-icon :name="changeIcon" :color="changeColor" size="16"></u-icon>
						{{ Math.abs(weightChange) }}kg {{ changeText }}
					</text>
				</view>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				showDatePicker: false,
				showHeightInput: false,

				weightForm: {
					record_date: Date.now(),
					weight: 0,
					height: 0
				},

				tempHeight: 170,
				weightString: '',
				previousWeight: 71.2
			}
		},
		computed: {
			weightDisplay() {
				return this.weightString || '0.0'
			},
			heightDisplay() {
				return this.weightForm.height > 0 ? `${this.weightForm.height} cm` : '未设置'
			},
			calculatedBMI() {
				if (this.weightForm.weight > 0 && this.weightForm.height > 0) {
					const heightInMeter = this.weightForm.height / 100
					return (this.weightForm.weight / (heightInMeter * heightInMeter)).toFixed(1)
				}
				return 0
			},
			bmiStatus() {
				const bmi = parseFloat(this.calculatedBMI)
				if (bmi <= 0) return ''
				if (bmi < 18.5) return '偏瘦'
				if (bmi < 24) return '正常'
				if (bmi < 28) return '超重'
				return '肥胖'
			},
			bmiStatusClass() {
				const bmi = parseFloat(this.calculatedBMI)
				if (bmi <= 0) return ''
				if (bmi < 18.5) return 'status-underweight'
				if (bmi < 24) return 'status-normal'
				if (bmi < 28) return 'status-overweight'
				return 'status-obese'
			},
			weightChange() {
				if (this.weightForm.weight <= 0 || this.previousWeight <= 0) return 0
				return this.weightForm.weight - this.previousWeight
			},
			changeText() {
				const change = this.weightChange
				if (change < 0) return '下降'
				if (change > 0) return '上升'
				return '无变化'
			},
			changeIcon() {
				const change = this.weightChange
				if (change < 0) return 'arrow-down'
				if (change > 0) return 'arrow-up'
				return 'minus'
			},
			changeColor() {
				const change = this.weightChange
				if (change < 0) return '#6bcf7f'
				if (change > 0) return '#ff6b6b'
				return '#999'
			},
			changeClass() {
				const change = this.weightChange
				if (change < 0) return 'change-down'
				if (change > 0) return 'change-up'
				return 'change-none'
			}
		},
		methods: {
			goBack() {
				uni.navigateBack()
			},

			saveWeight() {
				if (this.weightForm.weight <= 0) {
					uni.showToast({
						title: '请输入有效体重',
						icon: 'none'
					})
					return
				}

				// 这里应该调用API保存体重记录
				uni.showToast({
					title: '保存成功',
					icon: 'success'
				})

				// 返回上一页
				setTimeout(() => {
					uni.navigateBack()
				}, 1000)
			},

			confirmDate(e) {
				this.weightForm.record_date = e.value
				this.showDatePicker = false
			},

			keyboardChange(value) {
				// 限制最多输入5位数字（包括小数点）
				if (this.weightString.replace('.', '').length >= 5 && value !== '.') return

				// 如果已经有一个小数点，不能再输入小数点
				if (this.weightString.includes('.') && value === '.') return

				// 如果小数点后已经有1位数字，不能再输入
				const parts = this.weightString.split('.')
				if (parts.length > 1 && parts[1].length >= 1 && value !== '.' && this.weightString.includes('.')) return

				this.weightString += value
				this.updateWeightValue()
			},

			keyboardBackspace() {
				if (this.weightString.length > 0) {
					this.weightString = this.weightString.slice(0, -1)
					this.updateWeightValue()
				}
			},

			updateWeightValue() {
				const weight = parseFloat(this.weightString) || 0
				// 限制最大体重为999.9kg
				if (weight <= 999.9) {
					this.weightForm.weight = weight
				}
			},

			confirmHeight() {
				this.weightForm.height = this.tempHeight
				this.showHeightInput = false
			},

			formatDate(timestamp) {
				const date = new Date(timestamp)
				return `${date.getFullYear()}-${(date.getMonth() + 1).toString().padStart(2, '0')}-${date.getDate().toString().padStart(2, '0')}`
			}
		}
	}
</script>

<style lang="scss" scoped>
.container {
	padding: 20rpx;
}

.nav-btn {
	color: #3c9cff;
	font-size: 28rpx;
}

.form-container {
	.form-item {
		background: white;
		border-radius: 20rpx;
		padding: 0 30rpx 30rpx;
		margin-bottom: 20rpx;

		.label {
			display: block;
			font-size: 32rpx;
			font-weight: bold;
			color: #333;
			padding: 30rpx 0 20rpx;

			.optional {
				font-size: 24rpx;
				color: #999;
				font-weight: normal;
			}
		}

		.weight-input {
			.weight-display {
				text-align: center;
				padding: 50rpx 0;

				.weight-value {
					font-size: 80rpx;
					font-weight: bold;
					color: #333;
				}

				.unit {
					font-size: 32rpx;
					color: #999;
					margin-left: 10rpx;
				}
			}
		}

		.bmi-result {
			text-align: center;
			padding: 30rpx 0;

			.bmi-value {
				font-size: 60rpx;
				font-weight: bold;
				color: #3c9cff;
				margin-right: 20rpx;
			}

			.bmi-status {
				font-size: 28rpx;
				padding: 10rpx 20rpx;
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

		.bmi-reference {
			margin-top: 30rpx;
			padding-top: 30rpx;
			border-top: 1rpx solid #eee;

			.reference-title {
				font-size: 24rpx;
				color: #666;
				display: block;
				margin-bottom: 20rpx;
			}

			.reference-items {
				.reference-item {
					font-size: 24rpx;
					color: #999;
					display: block;
					margin-bottom: 10rpx;

					&:last-child {
						margin-bottom: 0;
					}

					.underweight {
						color: #d35400;
					}

					.normal {
						color: #55a630;
					}

					.overweight {
						color: #ff9f1c;
					}

					.obese {
						color: #e63946;
					}
				}
			}
		}

		.weight-change {
			display: flex;
			justify-content: space-between;
			align-items: center;
			padding: 30rpx 0;

			.change-value {
				display: flex;
				align-items: center;

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
		}
	}
}

.height-input-popup {
	padding: 30rpx;
	width: 600rpx;

	.popup-header {
		text-align: center;
		margin-bottom: 30rpx;

		.popup-title {
			font-size: 32rpx;
			font-weight: bold;
		}
	}

	.popup-buttons {
		display: flex;
		margin-top: 30rpx;
	}
}
</style>