<template>
	<view class="page" :style="{ paddingTop: topPadding + 'px' }">
		<view class="nav-bar">
			<text class="nav-back" @tap="goBack">‹ 返回</text>
			<text class="nav-title">版本变动</text>
			<view class="nav-placeholder"></view>
		</view>

		<!-- 当前版本卡片 -->
		<view class="current-card">
			<view class="current-badge"><text>当前版本</text></view>
			<view class="current-main">
				<text class="current-version">v{{ currentVersion }}</text>
				<text class="current-date">{{ changelog[0]?.date || '' }}</text>
			</view>
			<text class="current-title">{{ changelog[0]?.title || '' }}</text>
		</view>

		<!-- 版本时间线 -->
		<view class="timeline">
			<view v-for="(ver, idx) in changelog" :key="ver.version" class="timeline-item">
				<view class="timeline-rail">
					<view class="timeline-dot" :class="{ 'dot-current': idx === 0 }"></view>
					<view v-if="idx < changelog.length - 1" class="timeline-line"></view>
				</view>
				<view class="timeline-content">
					<view class="ver-card">
						<view class="ver-header">
							<view class="ver-header-left">
								<text class="ver-num">v{{ ver.version }}</text>
								<text class="ver-title">{{ ver.title }}</text>
							</view>
							<text class="ver-date">{{ ver.date }}</text>
						</view>
						<view class="change-list">
							<view v-for="(ch, i) in ver.changes" :key="i" class="change-item">
								<view class="change-tag" :class="'tag-' + ch.type">
									<text>{{ typeLabel(ch.type) }}</text>
								</view>
								<text class="change-text">{{ ch.text }}</text>
							</view>
						</view>
					</view>
				</view>
			</view>
		</view>

		<view class="footer-hint">
			<text>感谢你的陪伴 ❤️</text>
		</view>
	</view>
</template>

<script setup>
	import { computed } from 'vue'
	import { CHANGELOG, CURRENT_VERSION, TYPE_LABELS } from '../../utils/changelog'

	const changelog = CHANGELOG
	const currentVersion = CURRENT_VERSION

	const topPadding = computed(() => {
		const app = getApp()
		return (app.globalData?.customBarHeight || 88) + 8
	})

	const goBack = () => uni.navigateBack()

	const typeLabel = (type) => {
		return TYPE_LABELS[type]?.text || type
	}
</script>

<style lang="scss" scoped>
@import '../../styles/variables.scss';

.page {
	padding: 0 $spacing-xl $spacing-3xl;
	min-height: 100vh;
	background: $color-bg;
}

.nav-bar {
	@include nav-bar;
	.nav-back { font-size: $font-callout; color: $color-primary; font-weight: 500; }
	.nav-title { font-size: $font-headline; font-weight: 600; color: $color-label; }
	.nav-placeholder { width: 80rpx; }
}

// ---- Current Version Card ----
.current-card {
	@include card;
	padding: $spacing-xl;
	margin-bottom: $spacing-xl;
	background: linear-gradient(135deg, $color-primary 0%, $color-teal 100%);
	box-shadow: 0 8rpx 24rpx rgba(16, 185, 129, 0.2);
	position: relative;
	overflow: hidden;

	.current-badge {
		display: inline-block;
		padding: 4rpx $spacing-md;
		background: rgba(255, 255, 255, 0.22);
		border-radius: $radius-full;
		margin-bottom: $spacing-sm;

		text {
			font-size: $font-caption2;
			color: #fff;
			font-weight: 600;
			letter-spacing: 0.5rpx;
		}
	}

	.current-main {
		display: flex;
		align-items: baseline;
		gap: $spacing-md;
		margin-bottom: $spacing-xs;

		.current-version {
			font-size: 56rpx;
			font-weight: 700;
			color: #fff;
			letter-spacing: -1rpx;
			font-variant-numeric: tabular-nums;
		}
		.current-date {
			font-size: $font-caption1;
			color: rgba(255, 255, 255, 0.7);
			font-weight: 500;
			font-variant-numeric: tabular-nums;
		}
	}

	.current-title {
		display: block;
		font-size: $font-subhead;
		color: rgba(255, 255, 255, 0.88);
		font-weight: 500;
	}
}

// ---- Timeline ----
.timeline {
	display: flex;
	flex-direction: column;
}

.timeline-item {
	display: flex;
	align-items: stretch;
	gap: $spacing-md;
}

.timeline-rail {
	width: 24rpx;
	flex-shrink: 0;
	display: flex;
	flex-direction: column;
	align-items: center;
	padding-top: 28rpx;
}

.timeline-dot {
	width: 18rpx;
	height: 18rpx;
	border-radius: 50%;
	background: $color-separator-opaque;
	flex-shrink: 0;
	z-index: 1;

	&.dot-current {
		background: $color-primary;
		box-shadow: 0 0 0 6rpx rgba(16, 185, 129, 0.18);
	}
}

.timeline-line {
	flex: 1;
	width: 2rpx;
	background: $color-separator;
	margin-top: 4rpx;
}

.timeline-content {
	flex: 1;
	padding-bottom: $spacing-lg;
}

// ---- Version Card ----
.ver-card {
	@include card;
	padding: $spacing-lg $spacing-xl;

	.ver-header {
		display: flex;
		justify-content: space-between;
		align-items: baseline;
		margin-bottom: $spacing-md;
	}

	.ver-header-left {
		display: flex;
		align-items: baseline;
		gap: $spacing-sm;
		flex-wrap: wrap;
	}

	.ver-num {
		font-size: $font-headline;
		font-weight: 700;
		color: $color-label;
		font-variant-numeric: tabular-nums;
		letter-spacing: -0.5rpx;
	}

	.ver-title {
		font-size: $font-caption1;
		color: $color-label-tertiary;
		font-weight: 500;
	}

	.ver-date {
		font-size: $font-caption2;
		color: $color-label-quaternary;
		font-weight: 500;
		font-variant-numeric: tabular-nums;
		flex-shrink: 0;
	}
}

// ---- Change Items ----
.change-list {
	display: flex;
	flex-direction: column;
	gap: $spacing-sm;
}

.change-item {
	display: flex;
	align-items: flex-start;
	gap: $spacing-sm;
}

.change-tag {
	flex-shrink: 0;
	padding: 2rpx $spacing-xs;
	border-radius: $radius-sm;
	min-width: 56rpx;
	text-align: center;
	margin-top: 4rpx;

	text {
		font-size: $font-caption2;
		font-weight: 600;
		letter-spacing: 0.5rpx;
	}

	&.tag-feature {
		background: $color-primary-light;
		text { color: $color-primary; }
	}
	&.tag-improvement {
		background: $color-teal-light;
		text { color: $color-teal; }
	}
	&.tag-fix {
		background: $color-orange-light;
		text { color: $color-orange; }
	}
	&.tag-security {
		background: $color-red-light;
		text { color: $color-red; }
	}
}

.change-text {
	flex: 1;
	font-size: $font-subhead;
	color: $color-label;
	line-height: 1.5;
}

// ---- Footer ----
.footer-hint {
	text-align: center;
	padding: $spacing-xl 0 $spacing-md;

	text {
		font-size: $font-caption1;
		color: $color-label-quaternary;
		font-weight: 500;
	}
}
</style>
