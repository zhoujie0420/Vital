<template>
	<view class="page" :style="{ paddingTop: topPadding + 'px' }">
		<view class="nav-bar">
			<text class="nav-back" @tap="goBack">‹ 返回</text>
			<text class="nav-title">记录体重</text>
			<text class="nav-placeholder"></text>
		</view>

		<view class="card weight-card">
			<text class="card-label">体重</text>
			<view class="weight-display">
				<text class="weight-minus" @tap="adjustWeight(-0.1)">−</text>
				<view class="weight-center">
					<input type="digit" v-model="weight" class="weight-num" placeholder="0" />
					<text class="weight-unit">kg</text>
				</view>
				<text class="weight-plus" @tap="adjustWeight(0.1)">+</text>
			</view>
			<view class="quick-row">
				<text class="quick-pill" v-for="w in [50, 55, 60, 65, 70, 75, 80]" :key="w"
					:class="{ active: parseFloat(weight) === w }" @tap="weight = w">{{ w }}</text>
			</view>
		</view>

		<view class="card">
			<text class="card-label">身高（可选）</text>
			<view class="input-row">
				<input type="digit" v-model="height" class="field-input" placeholder="输入身高" />
				<text class="input-suffix">cm</text>
			</view>
		</view>

		<view class="card" v-if="bmi > 0">
			<view class="bmi-row">
				<view>
					<text class="bmi-label">BMI</text>
					<text class="bmi-status">{{ bmiStatus }}</text>
				</view>
				<text class="bmi-value">{{ bmi }}</text>
			</view>
		</view>

		<view class="bottom-action">
			<view class="save-btn" :class="{ loading: saving }" @tap="save">
				<text>{{ saving ? '保存中...' : '保存记录' }}</text>
			</view>
		</view>
	</view>
</template>

<script>
	import { createWeight } from '../../api/weight'

	export default {
		data() { return { saving: false, weight: '', height: '' } },
		computed: {
			topPadding() {
				const app = getApp()
				return (app.globalData?.customBarHeight || 88) + 8
			},
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
				if (c + d >= 0) this.weight = (c + d).toFixed(1)
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
.page {
	padding: 0 32rpx;
	padding-bottom: 160rpx;
	min-height: 100vh;
	background: #f2f2f7;
}

.nav-bar {
	display: flex; align-items: center; justify-content: space-between; margin-bottom: 28rpx;
	.nav-back { font-size: 32rpx; color: #007aff; font-weight: 500; }
	.nav-title { font-size: 34rpx; font-weight: 600; color: #1c1c1e; }
	.nav-placeholder { width: 80rpx; }
}

.card {
	background: #fff; border-radius: 20rpx; padding: 28rpx 32rpx; margin-bottom: 16rpx;
	box-shadow: 0 2rpx 16rpx rgba(0, 0, 0, 0.04);
	.card-label {
		display: block; font-size: 26rpx; font-weight: 600; color: #8e8e93;
		text-transform: uppercase; letter-spacing: 1rpx; margin-bottom: 20rpx;
	}
}

.weight-display {
	display: flex; align-items: center; justify-content: center; padding: 24rpx 0;

	.weight-minus, .weight-plus {
		width: 76rpx; height: 76rpx; line-height: 72rpx; text-align: center;
		background: #f2f2f7; border-radius: 50%; font-size: 44rpx; color: #007aff; font-weight: 300;
		transition: all 0.15s ease;
		&:active { background: #e5e5ea; transform: scale(0.88); }
	}

	.weight-center {
		display: flex; align-items: baseline; margin: 0 48rpx;
		.weight-num {
			width: 220rpx; text-align: center; font-size: 88rpx; font-weight: 700;
			color: #1c1c1e; letter-spacing: -3rpx;
		}
		.weight-unit { font-size: 32rpx; color: #8e8e93; margin-left: 4rpx; }
	}
}

.quick-row {
	display: flex; justify-content: center; flex-wrap: wrap; gap: 12rpx; margin-top: 16rpx;
	.quick-pill {
		padding: 12rpx 28rpx; background: #f2f2f7; border-radius: 20rpx;
		font-size: 26rpx; color: #636366; font-weight: 500;
		transition: all 0.2s ease;
		&.active { background: #007aff; color: #fff; box-shadow: 0 4rpx 12rpx rgba(0, 122, 255, 0.3); }
		&:active { transform: scale(0.92); }
	}
}

.input-row {
	display: flex; align-items: center; background: #f2f2f7; border-radius: 16rpx; padding: 22rpx;
	transition: box-shadow 0.2s ease;
	&:focus-within { box-shadow: 0 0 0 4rpx rgba(0, 122, 255, 0.15); }
	.field-input { flex: 1; font-size: 30rpx; color: #1c1c1e; }
	.input-suffix { font-size: 28rpx; color: #8e8e93; }
}

.bmi-row {
	display: flex; justify-content: space-between; align-items: center;
	.bmi-label { display: block; font-size: 30rpx; font-weight: 600; color: #1c1c1e; }
	.bmi-status {
		display: inline-block; font-size: 22rpx; color: #34c759; margin-top: 6rpx;
		background: rgba(52, 199, 89, 0.1); padding: 4rpx 16rpx; border-radius: 8rpx;
		font-weight: 600;
	}
	.bmi-value { font-size: 52rpx; font-weight: 700; color: #007aff; letter-spacing: -2rpx; }
}

.bottom-action {
	position: fixed; bottom: 0; left: 0; right: 0;
	padding: 20rpx 32rpx; padding-bottom: calc(20rpx + env(safe-area-inset-bottom));
	background: rgba(242, 242, 247, 0.85); backdrop-filter: blur(20px); -webkit-backdrop-filter: blur(20px);
	.save-btn {
		background: #007aff; color: #fff; text-align: center; padding: 28rpx 0;
		border-radius: 16rpx; font-size: 32rpx; font-weight: 600;
		box-shadow: 0 4rpx 16rpx rgba(0, 122, 255, 0.3);
		transition: all 0.15s ease;
		&:active { transform: scale(0.98); opacity: 0.9; }
		&.loading { opacity: 0.6; }
	}
}
</style>
