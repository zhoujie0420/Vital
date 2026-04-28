<template>
	<view class="page" :style="{ paddingTop: topPadding + 'px' }">
		<view class="nav-bar">
			<text class="nav-back" @tap="goBack">‹ 返回</text>
			<text class="nav-title">记录体重</text>
			<view class="nav-placeholder"></view>
		</view>

		<view class="card weight-card">
			<view class="weight-input-area" @tap="focusInput">
				<view class="weight-value-row">
					<text class="weight-value" :class="{ placeholder: !weight }">{{ weight || '0' }}</text>
					<text class="weight-unit">kg</text>
				</view>
				<text class="weight-hint">点击输入体重</text>
				<input
					ref="weightInput"
					type="digit"
					v-model="weight"
					class="hidden-input"
					:focus="inputFocused"
					@focus="inputFocused = true"
					@blur="inputFocused = false"
					placeholder=""
				/>
			</view>
		</view>

		<!-- 快捷选择 -->
		<view class="card">
			<text class="card-label">快捷选择</text>
			<view class="quick-grid">
				<text v-for="w in quickWeights" :key="w" class="quick-chip"
					:class="{ active: weight === String(w) }" @tap="weight = String(w)">
					{{ w }}
				</text>
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
		data() {
			return {
				saving: false,
				weight: '',
				inputFocused: false,
				quickWeights: [45, 50, 55, 60, 65, 70, 75, 80, 85, 90]
			}
		},
		computed: {
			topPadding() {
				const app = getApp()
				return (app.globalData?.customBarHeight || 88) + 8
			}
		},
		methods: {
			goBack() { uni.navigateBack() },
			focusInput() {
				this.inputFocused = true
			},
			async save() {
				const w = parseFloat(this.weight)
				if (!w || w <= 0) {
					uni.showToast({ title: '请输入体重', icon: 'none' })
					return
				}
				this.saving = true
				try {
					await createWeight({
						record_date: new Date().toISOString(),
						weight: w,
						height: 0
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

// --- Weight Input Area ---
.weight-card {
	padding: $spacing-3xl $spacing-xl;
}

.weight-input-area {
	text-align: center;
	position: relative;
}

.weight-value-row {
	display: flex;
	align-items: baseline;
	justify-content: center;
	margin-bottom: $spacing-sm;
}

.weight-value {
	font-size: 120rpx;
	font-weight: 700;
	color: $color-label;
	letter-spacing: -4rpx;
	font-variant-numeric: tabular-nums;
	line-height: 1;

	&.placeholder {
		color: $color-separator-opaque;
	}
}

.weight-unit {
	font-size: $font-title2;
	color: $color-label-quaternary;
	font-weight: 500;
	margin-left: $spacing-sm;
}

.weight-hint {
	display: block;
	font-size: $font-caption1;
	color: $color-label-quaternary;
	margin-top: $spacing-xs;
}

.hidden-input {
	position: absolute;
	left: 0;
	top: 0;
	width: 100%;
	height: 100%;
	opacity: 0;
}

// --- Quick Grid ---
.quick-grid {
	display: flex;
	flex-wrap: wrap;
	gap: $spacing-sm;

	.quick-chip {
		width: calc(20% - #{$spacing-sm});
		text-align: center;
		padding: $spacing-md 0;
		background: $color-fill;
		border-radius: $radius-md;
		font-size: $font-subhead;
		font-weight: 600;
		color: $color-label-tertiary;
		font-variant-numeric: tabular-nums;
		transition: all 0.2s cubic-bezier(0.16, 1, 0.3, 1);

		&.active {
			background: $color-primary;
			color: #fff;
			box-shadow: 0 4rpx 12rpx rgba(16, 185, 129, 0.2);
		}
		&:active { transform: scale(0.94); }
	}
}

// --- Bottom Action ---
.bottom-action {
	@include bottom-action-bar;
	.save-btn { @include primary-button; }
}
</style>
