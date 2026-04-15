<template>
	<view class="page" :style="{ paddingTop: topPadding + 'px' }">
		<view class="nav-bar">
			<text class="nav-back" @tap="goBack">‹ 返回</text>
			<text class="nav-title">设置</text>
			<text class="nav-placeholder"></text>
		</view>

		<view class="menu-group">
			<view class="menu-item">
				<view class="menu-left">
					<view class="menu-icon icon-purple"><text>📱</text></view>
					<text class="menu-text">版本</text>
				</view>
				<text class="menu-value">v1.0.0</text>
			</view>
			<view class="menu-item last" @tap="clearCache">
				<view class="menu-left">
					<view class="menu-icon icon-orange"><text>🗑️</text></view>
					<text class="menu-text">清除缓存</text>
				</view>
				<text class="arrow">›</text>
			</view>
		</view>
	</view>
</template>

<script setup>
	import { computed } from 'vue'

	const topPadding = computed(() => {
		const app = getApp()
		return (app.globalData?.customBarHeight || 88) + 8
	})

	const goBack = () => uni.navigateBack()
	const clearCache = () => {
		uni.showModal({
			title: '清除缓存', content: '确定清除？',
			success: (res) => {
				if (res.confirm) uni.showToast({ title: '已清除', icon: 'success' })
			}
		})
	}
</script>

<style lang="scss" scoped>
.page {
	padding: 0 32rpx;
	min-height: 100vh;
	background: #f2f2f7;
}

.nav-bar {
	display: flex;
	align-items: center;
	justify-content: space-between;
	margin-bottom: 28rpx;

	.nav-back { font-size: 32rpx; color: #007aff; font-weight: 500; }
	.nav-title { font-size: 34rpx; font-weight: 600; color: #1c1c1e; }
	.nav-placeholder { width: 80rpx; }
}

.menu-group {
	background: #fff;
	border-radius: 20rpx;
	overflow: hidden;
	box-shadow: 0 2rpx 16rpx rgba(0, 0, 0, 0.04);

	.menu-item {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: 24rpx 32rpx;
		border-bottom: 1rpx solid #f2f2f7;
		transition: background 0.15s ease;

		&:active { background: #f8f8f8; }
		&.last { border-bottom: none; }

		.menu-left {
			display: flex;
			align-items: center;
			gap: 20rpx;

			.menu-icon {
				width: 56rpx;
				height: 56rpx;
				border-radius: 14rpx;
				display: flex;
				align-items: center;
				justify-content: center;
				font-size: 28rpx;
			}
			.icon-purple { background: rgba(88, 86, 214, 0.1); }
			.icon-orange { background: rgba(255, 149, 0, 0.1); }

			.menu-text {
				font-size: 30rpx;
				color: #1c1c1e;
				font-weight: 500;
			}
		}

		.menu-value {
			font-size: 28rpx;
			color: #8e8e93;
		}
		.arrow {
			font-size: 36rpx;
			color: #c7c7cc;
			font-weight: 300;
		}
	}
}
</style>
