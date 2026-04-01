<template>
	<view class="container">
		<!-- 顶部装饰 -->
		<view class="header-decoration">
			<view class="decoration-circle circle-1"></view>
			<view class="decoration-circle circle-2"></view>
			<view class="decoration-circle circle-3"></view>
		</view>

		<!-- Logo和标题 -->
		<view class="logo-section">
			<view class="logo">
				<u-icon name="level" size="80" color="#3c9cff"></u-icon>
			</view>
			<text class="app-title">Vital Fitness</text>
			<text class="app-subtitle">记录每一个进步的瞬间</text>
		</view>

		<!-- 登录表单 -->
		<view class="login-form">
			<u-form :model="loginForm" ref="loginFormRef">
				<u-form-item>
					<u-input
						v-model="loginForm.username"
						placeholder="用户名/邮箱"
						prefix-icon="account"
						border="bottom"
						clearable
					/>
				</u-form-item>

				<u-form-item>
					<u-input
						v-model="loginForm.password"
						placeholder="密码"
						type="password"
						prefix-icon="lock"
						border="bottom"
						clearable
					/>
				</u-form-item>

				<view class="form-options">
					<view class="remember-me">
						<u-checkbox
							v-model="loginForm.rememberMe"
							active-color="#3c9cff"
						>
							<text class="checkbox-label">记住我</text>
						</u-checkbox>
					</view>
				</view>

				<u-button
					type="primary"
					@click="handleLogin"
					:loading="loading"
					:disabled="loading"
				>
					{{ loading ? '登录中...' : '登录' }}
				</u-button>
			</u-form>

			<!-- 注册入口 -->
			<view class="register-link">
				<text class="register-text">还没有账号？</text>
				<text class="register-btn" @click="toRegister">立即注册</text>
			</view>
		</view>
	</view>
</template>

<script>
	import { login } from '../../api/user'

	export default {
		data() {
			return {
				loading: false,
				loginForm: {
					username: '',
					password: '',
					rememberMe: false
				}
			}
		},
		onLoad() {
			// 如果记住了用户名，自动填充
			const savedUsername = uni.getStorageSync('remember_username')
			if (savedUsername) {
				this.loginForm.username = savedUsername
				this.loginForm.rememberMe = true
			}
		},
		methods: {
			async handleLogin() {
				if (!this.loginForm.username) {
					uni.showToast({ title: '请输入用户名', icon: 'none' })
					return
				}
				if (!this.loginForm.password || this.loginForm.password.length < 6) {
					uni.showToast({ title: '密码长度不能少于6位', icon: 'none' })
					return
				}

				this.loading = true
				try {
					const res = await login({
						username: this.loginForm.username,
						password: this.loginForm.password
					})

					// 保存登录信息到 store
					this.$store.dispatch('saveLoginInfo', {
						token: res.data.token,
						user: res.data.user
					})

					// 记住用户名
					if (this.loginForm.rememberMe) {
						uni.setStorageSync('remember_username', this.loginForm.username)
					} else {
						uni.removeStorageSync('remember_username')
					}

					uni.showToast({ title: '登录成功', icon: 'success' })
					setTimeout(() => {
						uni.switchTab({ url: '/pages/index/index' })
					}, 1000)
				} catch (err) {
					// request.js 已处理 toast
				} finally {
					this.loading = false
				}
			},

			toRegister() {
				uni.navigateTo({ url: '/pages/auth/register' })
			}
		}
	}
</script>

<style lang="scss" scoped>
.container {
	padding: 20rpx;
	min-height: 100vh;
	background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
	position: relative;
	overflow: hidden;
}

.header-decoration {
	.decoration-circle {
		position: absolute;
		border-radius: 50%;
		background: rgba(255, 255, 255, 0.1);

		&.circle-1 { width: 200rpx; height: 200rpx; top: -100rpx; right: -100rpx; }
		&.circle-2 { width: 300rpx; height: 300rpx; top: 200rpx; left: -150rpx; }
		&.circle-3 { width: 150rpx; height: 150rpx; bottom: 100rpx; right: 100rpx; }
	}
}

.logo-section {
	text-align: center;
	padding: 100rpx 0 80rpx;

	.logo { margin-bottom: 30rpx; }
	.app-title { font-size: 48rpx; font-weight: bold; color: white; display: block; margin-bottom: 20rpx; }
	.app-subtitle { font-size: 28rpx; color: rgba(255, 255, 255, 0.8); }
}

.login-form {
	background: white;
	border-radius: 20rpx;
	padding: 60rpx 40rpx;
	margin: 0 30rpx;

	.form-options {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin: 30rpx 0;

		.checkbox-label { font-size: 26rpx; color: #666; }
	}

	.register-link {
		text-align: center;
		margin-top: 40rpx;

		.register-text { font-size: 26rpx; color: #999; margin-right: 10rpx; }
		.register-btn { font-size: 26rpx; color: #3c9cff; font-weight: bold; }
	}
}
</style>
