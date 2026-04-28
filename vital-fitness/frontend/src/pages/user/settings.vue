<template>
	<view class="page" :style="{ paddingTop: topPadding + 'px' }">
		<view class="nav-bar">
			<text class="nav-back" @tap="goBack">‹ 返回</text>
			<text class="nav-title">设置</text>
			<view class="nav-placeholder"></view>
		</view>

		<view class="menu-card">
			<view class="menu-row">
				<view class="menu-left">
					<view class="menu-icon bg-purple">
						<text class="menu-icon-letter">V</text>
					</view>
					<text class="menu-text">版本</text>
				</view>
				<text class="menu-value">v1.0.0</text>
			</view>
			<view class="menu-row last" @tap="clearCache">
				<view class="menu-left">
					<view class="menu-icon bg-orange">
						<text class="menu-icon-letter">C</text>
					</view>
					<text class="menu-text">清除缓存</text>
				</view>
				<text class="sf-chevron">›</text>
			</view>
		</view>

		<view class="logout-btn" @tap="handleLogout">
			<text class="logout-text">退出登录</text>
		</view>
	</view>
</template>

<script setup>
	import { computed } from 'vue'
	import { logout } from '../../api/user'

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
	const handleLogout = () => {
		uni.showModal({
			title: '退出登录', content: '确定退出当前账号？',
			success: async (res) => {
				if (res.confirm) {
					try { await logout() } catch (e) {}
					uni.removeStorageSync('token')
					uni.removeStorageSync('userInfo')
					uni.redirectTo({ url: '/pages/auth/login' })
				}
			}
		})
	}
</script>

<style lang="scss" scoped>
@import '../../styles/variables.scss';

.page {
	padding: 0 $spacing-xl;
	min-height: 100vh;
	background: $color-bg;
}

.nav-bar {
	@include nav-bar;
	.nav-back { font-size: $font-callout; color: $color-primary; font-weight: 500; }
	.nav-title { font-size: $font-headline; font-weight: 600; color: $color-label; }
	.nav-placeholder { width: 80rpx; }
}

.menu-card {
	@include card;
	padding: 0;
	overflow: hidden;

	.menu-row {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: $spacing-lg $spacing-xl;
		position: relative;
		@include press-effect;

		&::after {
			content: '';
			position: absolute;
			bottom: 0;
			left: 100rpx;
			right: $spacing-xl;
			height: 0.5rpx;
			background: $color-separator;
		}

		&.last::after { display: none; }

		.menu-left {
			display: flex;
			align-items: center;
			gap: $spacing-md;

			.menu-icon {
				width: 60rpx;
				height: 60rpx;
				border-radius: $radius-sm;
				display: flex;
				align-items: center;
				justify-content: center;
				flex-shrink: 0;
			}
			.menu-icon-letter {
				font-size: $font-subhead;
				font-weight: 700;
			}
			.bg-purple {
				background: $color-purple-light;
				.menu-icon-letter { color: $color-purple; }
			}
			.bg-orange {
				background: $color-orange-light;
				.menu-icon-letter { color: $color-orange; }
			}

			.menu-text {
				font-size: $font-body;
				color: $color-label;
				font-weight: 400;
			}
		}

		.menu-value {
			font-size: $font-subhead;
			color: $color-label-quaternary;
		}
		.sf-chevron {
			font-size: 36rpx;
			color: $color-separator-opaque;
			font-weight: 300;
		}
	}
}

.logout-btn {
	margin-top: $spacing-2xl;
	@include card;
	padding: $spacing-lg 0;
	text-align: center;
	@include press-effect;

	.logout-text {
		font-size: $font-body;
		color: $color-red;
		font-weight: 500;
	}
}
</style>
