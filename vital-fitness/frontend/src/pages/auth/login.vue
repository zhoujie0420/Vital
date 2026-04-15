<template>
	<view class="page">
		<view class="bg-blur">
			<view class="bg-circle c1"></view>
			<view class="bg-circle c2"></view>
			<view class="bg-circle c3"></view>
		</view>

		<view class="content">
			<view class="logo-section">
				<view class="logo-ring">
					<text class="logo-emoji">💪</text>
				</view>
				<text class="app-title">Vital Fitness</text>
				<text class="app-subtitle">记录每一个进步的瞬间</text>
			</view>

			<view class="login-section">
				<view class="wx-login-btn" :class="{ loading: loading }" @tap="handleWxLogin">
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
.page {
	min-height: 100vh;
	background: #000;
	position: relative;
	overflow: hidden;
}

.bg-blur {
	position: absolute;
	top: 0; left: 0; right: 0; bottom: 0;

	.bg-circle {
		position: absolute;
		border-radius: 50%;
		filter: blur(80rpx);
	}
	.c1 {
		width: 500rpx; height: 500rpx;
		background: rgba(0, 122, 255, 0.5);
		top: -100rpx; right: -100rpx;
	}
	.c2 {
		width: 600rpx; height: 600rpx;
		background: rgba(88, 86, 214, 0.4);
		top: 300rpx; left: -200rpx;
	}
	.c3 {
		width: 400rpx; height: 400rpx;
		background: rgba(175, 82, 222, 0.3);
		bottom: 100rpx; right: -50rpx;
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
		background: rgba(255, 255, 255, 0.1);
		border: 3rpx solid rgba(255, 255, 255, 0.2);
		display: flex;
		align-items: center;
		justify-content: center;
		margin: 0 auto 32rpx;
		backdrop-filter: blur(20px);

		.logo-emoji {
			font-size: 72rpx;
		}
	}

	.app-title {
		display: block;
		font-size: 56rpx;
		font-weight: 700;
		color: #fff;
		letter-spacing: -1rpx;
		margin-bottom: 12rpx;
	}
	.app-subtitle {
		display: block;
		font-size: 28rpx;
		color: rgba(255, 255, 255, 0.6);
		font-weight: 400;
	}
}

.login-section {
	.wx-login-btn {
		background: rgba(255, 255, 255, 0.15);
		border: 1rpx solid rgba(255, 255, 255, 0.2);
		backdrop-filter: blur(20px);
		-webkit-backdrop-filter: blur(20px);
		border-radius: 20rpx;
		padding: 32rpx 0;
		text-align: center;
		transition: all 0.15s ease;

		&:active { transform: scale(0.98); background: rgba(255, 255, 255, 0.2); }
		&.loading { opacity: 0.6; }

		.btn-text {
			font-size: 32rpx;
			font-weight: 600;
			color: #fff;
		}
	}

	.login-hint {
		display: block;
		text-align: center;
		font-size: 24rpx;
		color: rgba(255, 255, 255, 0.35);
		margin-top: 20rpx;
	}
}
</style>
