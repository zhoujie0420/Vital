<template>
	<view class="container">
		<!-- 顶部导航栏 -->
		<u-navbar title="个人资料" fixed placeholder>
			<view slot="left">
				<u-icon name="arrow-left" size="20" @click="goBack"></u-icon>
			</view>
			<view slot="right">
				<text class="nav-btn" @click="saveProfile">保存</text>
			</view>
		</u-navbar>

		<view class="form-container">
			<!-- 头像 -->
			<view class="form-item avatar-item">
				<text class="label">头像</text>
				<view class="avatar-wrapper" @click="changeAvatar">
					<u-avatar
						:src="profileForm.avatar || '/static/images/default-avatar.png'"
						size="large"
					></u-avatar>
					<u-icon name="camera" size="20" color="#fff" class="camera-icon"></u-icon>
				</view>
			</view>

			<!-- 基本信息 -->
			<view class="form-item">
				<text class="label">基本信息</text>
				<u-cell-group>
					<u-cell title="昵称">
						<u-input
							slot="value"
							v-model="profileForm.nickname"
							placeholder="请输入昵称"
							border="none"
							style="text-align: right;"
						/>
					</u-cell>
					<u-cell title="用户名">
						<text slot="value" class="readonly-value">{{ profileForm.username }}</text>
					</u-cell>
					<u-cell title="性别">
						<u-picker
							:show="showGenderPicker"
							:columns="genderColumns"
							@confirm="confirmGender"
							@cancel="showGenderPicker = false"
						></u-picker>
						<u-cell
							slot="value"
							:title="getGenderText(profileForm.gender)"
							isLink
							@click="showGenderPicker = true"
						></u-cell>
					</u-cell>
					<u-cell title="生日">
						<u-datetime-picker
							:show="showBirthdayPicker"
							v-model="profileForm.birthday"
							mode="date"
							@confirm="confirmBirthday"
							@cancel="showBirthdayPicker = false"
						></u-datetime-picker>
						<u-cell
							slot="value"
							:title="formatDate(profileForm.birthday) || '未设置'"
							isLink
							@click="showBirthdayPicker = true"
						></u-cell>
					</u-cell>
				</u-cell-group>
			</view>

			<!-- 身体信息 -->
			<view class="form-item">
				<text class="label">身体信息</text>
				<u-cell-group>
					<u-cell title="身高">
						<view slot="value" class="number-input">
							<u-input
								v-model="profileForm.height"
								type="digit"
								placeholder="未设置"
								border="none"
								style="text-align: right; width: 100rpx;"
							/>
							<text class="unit">cm</text>
						</view>
					</u-cell>
					<u-cell title="体重">
						<view slot="value" class="number-input">
							<u-input
								v-model="profileForm.weight"
								type="digit"
								placeholder="未设置"
								border="none"
								style="text-align: right; width: 100rpx;"
							/>
							<text class="unit">kg</text>
						</view>
					</u-cell>
					<u-cell title="目标体重">
						<view slot="value" class="number-input">
							<u-input
								v-model="profileForm.target_weight"
								type="digit"
								placeholder="未设置"
								border="none"
								style="text-align: right; width: 100rpx;"
							/>
							<text class="unit">kg</text>
						</view>
					</u-cell>
				</u-cell-group>
			</view>

			<!-- 个人简介 -->
			<view class="form-item">
				<text class="label">个人简介</text>
				<u-textarea
					v-model="profileForm.bio"
					placeholder="介绍一下自己吧..."
					:height="120"
				/>
			</view>

			<!-- 联系方式 -->
			<view class="form-item">
				<text class="label">联系方式</text>
				<u-cell-group>
					<u-cell title="手机号">
						<text slot="value" class="readonly-value">{{ profileForm.phone || '未绑定' }}</text>
					</u-cell>
					<u-cell title="邮箱">
						<text slot="value" class="readonly-value">{{ profileForm.email || '未绑定' }}</text>
					</u-cell>
				</u-cell-group>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				showGenderPicker: false,
				showBirthdayPicker: false,

				profileForm: {
					avatar: '',
					username: 'fitness_user',
					nickname: '健身达人',
					gender: 1,
					birthday: null,
					height: '',
					weight: '',
					target_weight: '',
					bio: '坚持健身，享受生活',
					phone: '138****8888',
					email: 'user@example.com'
				},

				genderColumns: [['男', '女', '保密']]
			}
		},
		methods: {
			goBack() {
				uni.navigateBack()
			},

			saveProfile() {
				// 这里应该调用API保存用户资料
				uni.showToast({
					title: '保存成功',
					icon: 'success'
				})

				// 返回上一页
				setTimeout(() => {
					uni.navigateBack()
				}, 1000)
			},

			changeAvatar() {
				// 更换头像
				uni.chooseImage({
					count: 1,
					success: (res) => {
						// 这里应该上传头像到服务器
						this.profileForm.avatar = res.tempFilePaths[0]
						uni.showToast({
							title: '头像已选择',
							icon: 'success'
						})
					}
				})
			},

			confirmGender(e) {
				// 0: 保密, 1: 男, 2: 女
				const genderMap = {
					'男': 1,
					'女': 2,
					'保密': 0
				}
				this.profileForm.gender = genderMap[e.value[0]] || 0
				this.showGenderPicker = false
			},

			confirmBirthday(e) {
				this.profileForm.birthday = e.value
				this.showBirthdayPicker = false
			},

			getGenderText(gender) {
				const genderText = {
					0: '保密',
					1: '男',
					2: '女'
				}
				return genderText[gender] || '未设置'
			},

			formatDate(timestamp) {
				if (!timestamp) return ''
				const date = new Date(timestamp)
				return `${date.getFullYear()}-${(date.getMonth() + 1).toString().padStart(2, '0')}-${date.getDate().toString().padStart(2, '0')}`
			}
		}
	}
</script>

<style lang="scss" scoped>
.container {
	padding: 20rpx;
}

.nav-btn {
	color: #3c9cff;
	font-size: 28rpx;
}

.form-container {
	.form-item {
		background: white;
		border-radius: 20rpx;
		padding: 0 30rpx 30rpx;
		margin-bottom: 20rpx;

		&.avatar-item {
			text-align: center;
			padding: 40rpx 30rpx;
		}

		.label {
			display: block;
			font-size: 32rpx;
			font-weight: bold;
			color: #333;
			padding: 30rpx 0 20rpx;
		}

		.avatar-wrapper {
			position: relative;
			display: inline-block;

			.camera-icon {
				position: absolute;
				bottom: 0;
				right: 0;
				background: #3c9cff;
				border-radius: 50%;
				padding: 5rpx;
			}
		}

		.readonly-value {
			color: #999;
			font-size: 28rpx;
		}

		.number-input {
			display: flex;
			align-items: center;
			justify-content: flex-end;

			.unit {
				font-size: 28rpx;
				color: #999;
				margin-left: 10rpx;
			}
		}
	}
}
</style>