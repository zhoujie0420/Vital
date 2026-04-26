<template>
	<view class="page">
		<view class="bg-layer">
			<view class="bg-orb orb-1"></view>
			<view class="bg-orb orb-2"></view>
			<view class="bg-orb orb-3"></view>
		</view>

		<view class="content">
			<view class="logo-section">
				<view class="logo-ring">
					<view class="logo-ring-inner">
						<text class="logo-emoji">💪</text>
					</view>
				</view>
				<text class="app-title">Vital Fitness</text>
				<text class="app-subtitle">记录每一个进步的瞬间</text>
			</view>

			<view class="login-section">
				<view class="login-btn" :class="{ loading: loading }" @tap="handleWxLogin">
					<text class="btn-text">{{ loading ? '登录中...' : '微信一键登录' }}</text>
				</view>
				<text class="login-hint">使用微信账号安全登录</text>
			</view>
		</view>
	</view>
</template>

<script setup>
	import { ref } from 'vue'
	import { useUserStore } from '../../store'
	import { wxLogin } from '../../api/user'

	const loading = ref(false)
	const userStore = useUserStore()

	const handleWxLogin = () => {
		loading.value = true
		uni.login({
			provider: 'weixin',
			success: async (loginRes) => {
				try {
					const res = await wxLogin(loginRes.code)
					userStore.saveLoginInfo({ token: res.data.token, user: res.data.user })
					uni.showToast({ title: '登录成功', icon: 'success' })
					setTimeout(() => uni.switchTab({ url: '/pages/index/index' }), 1000)
				} catch (err) {
					uni.showToast({ title: '登录失败', icon: 'none' })
				} finally { loading.value = false }
			},
			fail: () => {
				loading.value = false
				uni.showToast({ title: '微信登录取消', icon: 'none' })
			}
		})
	}
</script>

<style lang="scss" scoped>
@import '../../styles/variables.scss';

.page {
	min-height: 100vh;
	background: #000;
	position: relative;
	overflow: hidden;
}

.bg-layer {
	position: absolute;
	top: 0; left: 0; right: 0; bottom: 0;

	.bg-orb {
		position: absolute;
		border-radius: 50%;
		filter: blur(100rpx);
	}
	.orb-1 {
		width: 480rpx; height: 480rpx;
		background: rgba(0, 122, 255, 0.45);
		top: -80rpx; right: -80rpx;
	}
	.orb-2 {
		width: 560rpx; height: 560rpx;
		background: rgba(88, 86, 214, 0.35);
		top: 320rpx; left: -180rpx;
	}
	.orb-3 {
		width: 380rpx; height: 380rpx;
		background: rgba(175, 82, 222, 0.25);
		bottom: 120rpx; right: -40rpx;
	}
}

.content {
	position: relative;
	z-index: 1;
	min-height: 100vh;
	display: flex;
	flex-direction: column;
	justify-content: center;
	padding: 60rpx 48rpx;
}

.logo-section {
	text-align: center;
	margin-bottom: 120rpx;

	.logo-ring {
		width: 160rpx;
		height: 160rpx;
		border-radius: 50%;
		border: 3rpx solid rgba(255, 255, 255, 0.12);
		display: flex;
		align-items: center;
		justify-content: center;
		margin: 0 auto 32rpx;

		.logo-ring-inner {
			width: 130rpx;
			height: 130rpx;
			border-radius: 50%;
			background: rgba(255, 255, 255, 0.06);
			backdrop-filter: blur(20px);
			-webkit-backdrop-filter: blur(20px);
			display: flex;
			align-items: center;
			justify-content: center;
		}

		.logo-emoji {
			font-size: 64rpx;
		}
	}

	.app-title {
		display: block;
		font-size: 56rpx;
		font-weight: 700;
		color: #fff;
		letter-spacing: -1rpx;
		margin-bottom: $spacing-sm;
	}
	.app-subtitle {
		display: block;
		font-size: $font-subhead;
		color: rgba(255, 255, 255, 0.45);
		font-weight: 400;
	}
}

.login-section {
	.login-btn {
		background: rgba(255, 255, 255, 0.12);
		border: 0.5rpx solid rgba(255, 255, 255, 0.15);
		backdrop-filter: saturate(180%) blur(40px);
		-webkit-backdrop-filter: saturate(180%) blur(40px);
		border-radius: $radius-lg;
		padding: $spacing-xl 0;
		text-align: center;
		transition: all 0.2s cubic-bezier(0.25, 0.46, 0.45, 0.94);

		&:active { transform: scale(0.98); background: rgba(255, 255, 255, 0.18); }
		&.loading { opacity: 0.5; pointer-events: none; }

		.btn-text {
			font-size: $font-callout;
			font-weight: 600;
			color: #fff;
		}
	}

	.login-hint {
		display: block;
		text-align: center;
		font-size: $font-caption1;
		color: rgba(255, 255, 255, 0.25);
		margin-top: $spacing-lg;
	}
}
</style>
