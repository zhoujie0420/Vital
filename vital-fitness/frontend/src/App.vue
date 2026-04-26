<script>
	export default {
		onLaunch() {
			const token = uni.getStorageSync('token')
			if (!token) {
				uni.redirectTo({ url: '/pages/auth/login' })
			}
			// 获取系统信息，设置全局状态栏高度
			const sysInfo = uni.getSystemInfoSync()
			const statusBarHeight = sysInfo.statusBarHeight || 20
			let navBarHeight = 44
			let customBarHeight = statusBarHeight + navBarHeight
			try {
				const menuBtn = uni.getMenuButtonBoundingClientRect()
				if (menuBtn && menuBtn.top > 0) {
					navBarHeight = (menuBtn.top - statusBarHeight) * 2 + menuBtn.height
					customBarHeight = statusBarHeight + navBarHeight
				}
			} catch (e) {}
			this.globalData = {
				statusBarHeight,
				navBarHeight,
				customBarHeight
			}
		},
		globalData: {
			statusBarHeight: 44,
			navBarHeight: 44,
			customBarHeight: 88
		}
	}
</script>

<style lang="scss">
	@import './styles/variables.scss';

	page {
		background-color: $color-bg;
		font-family: -apple-system, BlinkMacSystemFont, 'SF Pro Display', 'SF Pro Text',
			'Helvetica Neue', 'PingFang SC', sans-serif;
		-webkit-font-smoothing: antialiased;
		-moz-osx-font-smoothing: grayscale;
		color: $color-label;
		line-height: 1.4;
	}

	/* 全局滚动条隐藏 */
	::-webkit-scrollbar {
		display: none;
		width: 0;
		height: 0;
	}

	/* 全局 view 盒模型 */
	view, text {
		box-sizing: border-box;
	}
</style>
