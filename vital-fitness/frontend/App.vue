<script>
	export default {
		onLaunch() {
			console.log('App Launch')
			this.checkLoginStatus()
		},
		onShow() {
			console.log('App Show')
		},
		onHide() {
			console.log('App Hide')
		},
		methods: {
			checkLoginStatus() {
				const token = uni.getStorageSync('token')
				if (!token) {
					// 未登录，跳转到登录页
					uni.redirectTo({ url: '/pages/auth/login' })
				}
			}
		}
	}
</script>

<style lang="scss">
	@import "uview-ui/index.scss";

	page {
		background-color: #f5f5f5;
	}
</style>
