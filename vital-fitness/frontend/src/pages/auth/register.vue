<template>
	<view class="page">
		<view class="content">
			<text class="hint">请使用微信登录</text>
			<view class="link-btn" @tap="toLogin">
				<text class="link-text">去登录</text>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		methods: {
			toLogin() {
				uni.redirectTo({ url: '/pages/auth/login' })
			}
		}
	}
</script>

<style lang="scss" scoped>
@import '../../styles/variables.scss';

.page {
	min-height: 100vh;
	background: $color-bg;
	display: flex;
	align-items: center;
	justify-content: center;
}

.content {
	text-align: center;
	padding: $spacing-2xl;

	.hint {
		display: block;
		font-size: $font-body;
		color: $color-label-quaternary;
		margin-bottom: $spacing-xl;
	}

	.link-btn {
		display: inline-block;
		padding: $spacing-md $spacing-3xl;
		@include primary-button;
	}
}
</style>
