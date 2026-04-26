<template>
	<view class="page" :style="{ paddingTop: topPadding + 'px' }">
		<view class="nav-bar">
			<text class="nav-back" @tap="goBack">‹ 返回</text>
			<text class="nav-title">记录体重</text>
			<view class="nav-placeholder"></view>
		</view>

		<view class="card weight-card">
			<text class="card-label">体重</text>
			<view class="weight-display">
				<view class="weight-btn" @tap="adjustWeight(-0.1)">
					<text class="weight-btn-text">−</text>
				</view>
				<view class="weight-center">
					<input type="digit" v-model="weight" class="weight-num" placeholder="0" />
					<text class="weight-unit">kg</text>
				</view>
				<view class="weight-btn" @tap="adjustWeight(0.1)">
					<text class="weight-btn-text">+</text>
				</view>
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

		<view class="card bmi-card" v-if="bmi > 0">
			<view class="bmi-row">
				<view class="bmi-left">
					<text class="bmi-label">BMI</text>
					<view class="bmi-badge" :class="bmiClass">
						<text>{{ bmiStatus }}</text>
					</view>
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
			},
			bmiClass() {
				const b = parseFloat(this.bmi)
				if (b < 18.5) return 'bmi-thin'
				if (b < 24) return 'bmi-normal'
				if (b < 28) return 'bmi-over'
				return 'bmi-obese'
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
@import '../../styles/variables.scss';

.page {
	padding: 0 $spacing-xl;
	padding-bottom: 160rpx;
	min-height: 100vh;
	background: $color-bg;
}

.nav-bar {
	@include nav-bar;
	.nav-back { font-size: $font-callout; color: $color-primary; font-weight: 500; }
	.nav-title { font-size: $font-headline; font-weight: 600; color: $color-label; }
	.nav-placeholder { width: 80rpx; }
}

.card {
	@include card;
	margin-bottom: $spacing-md;
	.card-label { @include card-label; }
}

// --- Weight Display ---
.weight-display {
	display: flex;
	align-items: center;
	justify-content: center;
	padding: $spacing-lg 0;

	.weight-btn {
		width: 76rpx;
		height: 76rpx;
		border-radius: 50%;
		background: $color-fill;
		display: flex;
		align-items: center;
		justify-content: center;
		@include press-effect;

		.weight-btn-text {
			font-size: 44rpx;
			color: $color-primary;
			font-weight: 300;
			line-height: 1;
		}
	}

	.weight-center {
		display: flex;
		align-items: baseline;
		margin: 0 $spacing-3xl;

		.weight-num {
			width: 240rpx;
			text-align: center;
			font-size: 96rpx;
			font-weight: 700;
			color: $color-label;
			letter-spacing: -3rpx;
			font-variant-numeric: tabular-nums;
		}
		.weight-unit {
			font-size: $font-callout;
			color: $color-label-quaternary;
			margin-left: 4rpx;
		}
	}
}

.quick-row {
	display: flex;
	justify-content: center;
	flex-wrap: wrap;
	gap: $spacing-sm;
	margin-top: $spacing-md;

	.quick-pill {
		padding: $spacing-sm $spacing-lg;
		background: $color-fill;
		border-radius: $radius-full;
		font-size: $font-footnote;
		color: $color-label-tertiary;
		font-weight: 500;
		font-variant-numeric: tabular-nums;
		transition: all 0.25s cubic-bezier(0.25, 0.46, 0.45, 0.94);

		&.active { background: $color-label; color: #fff; }
		&:active { transform: scale(0.92); }
	}
}

// --- Input ---
.input-row {
	@include input-field;
	.field-input { flex: 1; font-size: $font-body; color: $color-label; }
	.input-suffix { font-size: $font-subhead; color: $color-label-quaternary; }
}

// --- BMI ---
.bmi-row {
	display: flex;
	justify-content: space-between;
	align-items: center;

	.bmi-left {
		.bmi-label {
			display: block;
			font-size: $font-body;
			font-weight: 600;
			color: $color-label;
			margin-bottom: $spacing-xs;
		}
		.bmi-badge {
			display: inline-block;
			font-size: $font-caption2;
			padding: 4rpx $spacing-md;
			border-radius: $spacing-xs;
			font-weight: 600;
		}
		.bmi-thin { background: $color-primary-light; color: $color-primary; }
		.bmi-normal { background: $color-green-light; color: $color-green; }
		.bmi-over { background: $color-orange-light; color: $color-orange; }
		.bmi-obese { background: $color-red-light; color: $color-red; }
	}

	.bmi-value {
		font-size: 56rpx;
		font-weight: 700;
		color: $color-primary;
		letter-spacing: -2rpx;
		font-variant-numeric: tabular-nums;
	}
}

.bottom-action {
	@include bottom-action-bar;
	.save-btn { @include primary-button; }
}
</style>
