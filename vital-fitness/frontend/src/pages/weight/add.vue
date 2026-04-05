<template>
	<view class="container">
		<u-navbar title="记录体重" fixed placeholder>
			<view slot="left">
				<u-icon name="arrow-left" size="20" @click="goBack"></u-icon>
			</view>
		</u-navbar>

		<view class="form-container">
			<view class="form-item">
				<text class="label">体重 (kg)</text>
				<view class="weight-input">
					<text class="weight-btn" @tap="adjustWeight(-0.5)">-</text>
					<input type="digit" v-model="weight" class="weight-value" placeholder="0" />
					<text class="weight-btn" @tap="adjustWeight(0.5)">+</text>
				</view>
				<view class="quick-weights">
					<text class="quick-btn" v-for="w in [50, 60, 70, 80, 90]" :key="w" @tap="weight = w">{{w}}</text>
				</view>
			</view>

			<view class="form-item">
				<text class="label">身高 (cm) <text class="optional">可选</text></text>
				<input type="digit" v-model="height" class="input-field" placeholder="输入身高" />
			</view>

			<view class="form-item" v-if="bmi > 0">
				<text class="label">BMI</text>
				<text class="bmi-value">{{ bmi }} - {{ bmiStatus }}</text>
			</view>

			<button class="save-btn" type="primary" :loading="saving" @tap="save">
				{{ saving ? '保存中...' : '保存记录' }}
			</button>
		</view>
	</view>
</template>

<script>
	import { createWeight } from '../../api/weight'

	export default {
		data() {
			return { saving: false, weight: '', height: '' }
		},
		computed: {
			bmi() {
				const w = parseFloat(this.weight), h = parseFloat(this.height)
				if (w > 0 && h > 0) return (w / ((h / 100) * (h / 100))).toFixed(1)
				return 0
			},
			bmiStatus() {
				const b = parseFloat(this.bmi)
				if (b < 18.5) return '偏瘦'
				if (b < 24) return '正常'
				if (b < 28) return '超重'
				return '肥胖'
			}
		},
		methods: {
			goBack() { uni.navigateBack() },
			adjustWeight(d) {
				const c = parseFloat(this.weight) || 0
				if (c + d >= 0) this.weight = c + d
			},
			async save() {
				if (!this.weight || parseFloat(this.weight) <= 0) {
					uni.showToast({ title: '请输入体重', icon: 'none' })
					return
				}
				this.saving = true
				try {
					await createWeight({
						record_date: new Date().toISOString(),
						weight: parseFloat(this.weight),
						height: parseFloat(this.height) || 0
					})
					uni.showToast({ title: '保存成功', icon: 'success' })
					setTimeout(() => uni.navigateBack(), 1000)
				} catch (e) {} finally { this.saving = false }
			}
		}
	}
</script>

<style lang="scss" scoped>
.container { padding: 20rpx; }
.form-container {
	.form-item {
		background: white; border-radius: 20rpx; padding: 30rpx; margin-bottom: 20rpx;
		.label { font-size: 32rpx; font-weight: bold; color: #333; margin-bottom: 20rpx; display: block;
			.optional { font-size: 24rpx; color: #999; font-weight: normal; }
		}
		.weight-input {
			display: flex; align-items: center; justify-content: center; margin: 20rpx 0;
			.weight-btn { width: 80rpx; height: 80rpx; line-height: 80rpx; text-align: center; background: #3c9cff; color: white; border-radius: 50%; font-size: 40rpx; }
			.weight-value { width: 200rpx; text-align: center; font-size: 60rpx; font-weight: bold; margin: 0 30rpx; }
		}
		.quick-weights { display: flex; justify-content: center; gap: 20rpx;
			.quick-btn { padding: 10rpx 24rpx; background: #f5f5f5; border-radius: 10rpx; font-size: 24rpx; color: #666; }
		}
		.input-field { border: 1rpx solid #eee; border-radius: 10rpx; padding: 20rpx; font-size: 28rpx; }
		.bmi-value { font-size: 36rpx; font-weight: bold; color: #3c9cff; }
	}
	.save-btn { margin-top: 20rpx; border-radius: 15rpx; font-size: 32rpx; }
}
</style>
