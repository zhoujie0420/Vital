<template>
	<view class="container">
		<!-- 顶部导航栏 -->
		<u-navbar title="通用设置" fixed placeholder>
			<view slot="left">
				<u-icon name="arrow-left" size="20" @click="goBack"></u-icon>
			</view>
		</u-navbar>

		<view class="settings-container">
			<!-- 显示设置 -->
			<view class="settings-section">
				<u-cell-group>
					<u-cell title="深色模式">
						<u-switch
							slot="right-icon"
							v-model="settings.darkMode"
							active-color="#3c9cff"
						></u-switch>
					</u-cell>
					<u-cell title="显示单位">
						<u-picker
							:show="showUnitPicker"
							:columns="unitColumns"
							@confirm="confirmUnit"
							@cancel="showUnitPicker = false"
						></u-picker>
						<u-cell
							slot="value"
							:title="settings.unit"
							isLink
							@click="showUnitPicker = true"
						></u-cell>
					</u-cell>
					<u-cell title="首页卡片">
						<u-switch
							slot="right-icon"
							v-model="settings.showHomeCards"
							active-color="#3c9cff"
						></u-switch>
					</u-cell>
				</u-cell-group>
			</view>

			<!-- 提醒设置 -->
			<view class="settings-section">
				<u-cell-group>
					<u-cell title="训练提醒">
						<u-switch
							slot="right-icon"
							v-model="settings.workoutReminders"
							active-color="#3c9cff"
						></u-switch>
					</u-cell>
					<u-cell
						title="提醒时间"
						:value="settings.reminderTime"
						isLink
						@click="setReminderTime"
						:disabled="!settings.workoutReminders"
					></u-cell>
					<u-cell title="饮食提醒">
						<u-switch
							slot="right-icon"
							v-model="settings.dietReminders"
							active-color="#3c9cff"
						></u-switch>
					</u-cell>
					<u-cell title="体重记录提醒">
						<u-switch
							slot="right-icon"
							v-model="settings.weightReminders"
							active-color="#3c9cff"
						></u-switch>
					</u-cell>
				</u-cell-group>
			</view>

			<!-- 数据设置 -->
			<view class="settings-section">
				<u-cell-group>
					<u-cell title="自动同步">
						<u-switch
							slot="right-icon"
							v-model="settings.autoSync"
							active-color="#3c9cff"
						></u-switch>
					</u-cell>
					<u-cell title="WiFi下自动同步">
						<u-switch
							slot="right-icon"
							v-model="settings.wifiSync"
							active-color="#3c9cff"
							:disabled="!settings.autoSync"
						></u-switch>
					</u-cell>
					<u-cell
						title="清除缓存"
						isLink
						@click="clearCache"
					></u-cell>
					<u-cell
						title="导出数据"
						isLink
						@click="exportData"
					></u-cell>
				</u-cell-group>
			</view>

			<!-- 关于 -->
			<view class="settings-section">
				<u-cell-group>
					<u-cell
						title="检查更新"
						:value="`v${appVersion}`"
						isLink
						@click="checkUpdate"
					></u-cell>
					<u-cell
						title="功能介绍"
						isLink
						@click="showFeatures"
					></u-cell>
				</u-cell-group>
			</view>

			<!-- 操作按钮 -->
			<view class="action-buttons">
				<u-button
					type="primary"
					plain
					@click="backupSettings"
				>备份设置</u-button>
				<u-button
					type="primary"
					plain
					@click="restoreSettings"
				>恢复设置</u-button>
			</view>
		</view>

		<!-- 时间选择器 -->
		<u-datetime-picker
			:show="showTimePicker"
			v-model="tempReminderTime"
			mode="time"
			@confirm="confirmReminderTime"
			@cancel="showTimePicker = false"
		></u-datetime-picker>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				showUnitPicker: false,
				showTimePicker: false,

				appVersion: '1.0.0',
				tempReminderTime: Date.now(),

				settings: {
					darkMode: false,
					unit: '公制(kg/cm)',
					showHomeCards: true,
					workoutReminders: true,
					reminderTime: '18:00',
					dietReminders: true,
					weightReminders: true,
					autoSync: true,
					wifiSync: true
				},

				unitColumns: [[
					'公制(kg/cm)',
					'英制(lbs/in)'
				]]
			}
		},
		methods: {
			goBack() {
				uni.navigateBack()
			},

			confirmUnit(e) {
				this.settings.unit = e.value[0]
				this.showUnitPicker = false
			},

			setReminderTime() {
				if (!this.settings.workoutReminders) return
				this.showTimePicker = true
			},

			confirmReminderTime(e) {
				const date = new Date(e.value)
				this.settings.reminderTime = `${date.getHours().toString().padStart(2, '0')}:${date.getMinutes().toString().padStart(2, '0')}`
				this.showTimePicker = false
			},

			clearCache() {
				uni.showModal({
					title: '清除缓存',
					content: '确定要清除所有缓存数据吗？这不会影响您的个人记录。',
					success: (res) => {
						if (res.confirm) {
							// 这里应该执行清除缓存的逻辑
							uni.showToast({
								title: '缓存已清除',
								icon: 'success'
							})
						}
					}
				})
			},

			exportData() {
				uni.showActionSheet({
					itemList: ['导出为Excel', '导出为CSV', '导出为PDF'],
					success: (res) => {
						uni.showToast({
							title: `正在导出为${['Excel', 'CSV', 'PDF'][res.tapIndex]}`,
							icon: 'success'
						})
					}
				})
			},

			checkUpdate() {
				// 检查更新逻辑
				uni.showToast({
					title: '当前已是最新版本',
					icon: 'success'
				})
			},

			showFeatures() {
				// 显示功能介绍
				uni.showToast({
					title: '功能介绍',
					icon: 'none'
				})
			},

			backupSettings() {
				// 备份设置
				uni.setStorageSync('backup_settings', this.settings)
				uni.showToast({
					title: '设置已备份',
					icon: 'success'
				})
			},

			restoreSettings() {
				// 恢复设置
				const backup = uni.getStorageSync('backup_settings')
				if (backup) {
					this.settings = { ...this.settings, ...backup }
					uni.showToast({
						title: '设置已恢复',
						icon: 'success'
					})
				} else {
					uni.showToast({
						title: '暂无备份设置',
						icon: 'none'
					})
				}
			}
		}
	}
</script>

<style lang="scss" scoped>
.container {
	padding: 20rpx;
}

.settings-container {
	.settings-section {
		background: white;
		border-radius: 20rpx;
		margin-bottom: 20rpx;
		overflow: hidden;
	}

	.action-buttons {
		display: flex;
		gap: 20rpx;
		padding: 30rpx 0;

		::v-deep .u-btn {
			flex: 1;
			border-radius: 10rpx;
		}
	}
}
</style>