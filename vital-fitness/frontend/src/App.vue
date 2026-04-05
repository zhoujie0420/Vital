<script>
	export default {
		onLaunch() {
			const token = uni.getStorageSync('token')
			if (!token) {
				uni.redirectTo({ url: '/pages/auth/login' })
			}
		}
	}
</script>

<style>
	page { background-color: #f2f2f7; }
</style>
