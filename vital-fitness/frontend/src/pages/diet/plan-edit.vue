<template>
	<view class="page" :style="{ paddingTop: topPadding + 'px' }">
		<view class="nav-bar">
			<text class="nav-back" @tap="goBack">‹ 返回</text>
			<text class="nav-title">{{ pageTitle }}</text>
			<view class="nav-placeholder"></view>
		</view>

		<!-- 步骤1：选择预设模板（仅 mode=system） -->
		<view v-if="mode === 'system' && !templateId" class="card">
			<text class="card-label">选择方案类型</text>
			<view v-for="tpl in systemTemplates" :key="tpl.id" class="preset-card"
				:class="{ selected: selectedPreset === tpl.id }" @tap="selectPreset(tpl)">
				<view class="preset-top">
					<text class="preset-name">{{ tpl.name }}</text>
					<text class="preset-type">{{ typeLabel(tpl.type) }}</text>
				</view>
				<text class="preset-desc">{{ tpl.description }}</text>
			</view>
		</view>

		<!-- 步骤2：填写方案参数（选完模板后 或 custom/edit 模式） -->
		<template v-if="templateId || mode === 'custom' || mode === 'edit'">
			<!-- 方案名称 -->
			<view class="card">
				<text class="card-label">方案名称</text>
				<view class="input-row">
					<input v-model="form.name" placeholder="输入方案名称" class="field-input" />
				</view>
			</view>

			<!-- 固定比例 / 自定义 宏量输入 -->
		<view class="card" v-if="form.type === 'fixed' || mode === 'custom'">
			<text class="card-label">每日营养目标</text>
			<view class="macro-grid">
				<view class="macro-field">
					<text class="macro-label">热量</text>
					<view class="macro-input-wrap">
						<input type="digit" v-model="form.calories" class="macro-input" placeholder="0" />
						<text class="macro-suffix">kcal</text>
					</view>
					<text v-if="errors.calories" class="field-error">{{ errors.calories }}</text>
				</view>
				<view class="macro-field">
					<text class="macro-label">蛋白质</text>
					<view class="macro-input-wrap">
						<input type="digit" v-model="form.protein" class="macro-input" placeholder="0" />
						<text class="macro-suffix">g</text>
					</view>
					<text v-if="errors.protein" class="field-error">{{ errors.protein }}</text>
				</view>
				<view class="macro-field">
					<text class="macro-label">碳水</text>
					<view class="macro-input-wrap">
						<input type="digit" v-model="form.carbs" class="macro-input" placeholder="0" />
						<text class="macro-suffix">g</text>
					</view>
					<text v-if="errors.carbs" class="field-error">{{ errors.carbs }}</text>
				</view>
				<view class="macro-field">
					<text class="macro-label">脂肪</text>
					<view class="macro-input-wrap">
						<input type="digit" v-model="form.fat" class="macro-input" placeholder="0" />
						<text class="macro-suffix">g</text>
					</view>
					<text v-if="errors.fat" class="field-error">{{ errors.fat }}</text>
				</view>
			</view>
		</view>

		<!-- 碳循环配置 -->
		<view class="card" v-if="form.type === 'carb_cycle'">
			<text class="card-label">碳循环 · 周配置</text>
			<view class="cycle-week">
				<view v-for="(day, i) in weekDays" :key="i" class="cycle-day"
					:class="{ high: form.dayMap[i] === 'high', low: form.dayMap[i] === 'low' }"
					@tap="toggleDay(i)">
					<text class="cycle-day-name">{{ day }}</text>
					<text class="cycle-day-type">{{ form.dayMap[i] === 'high' ? '高碳' : '低碳' }}</text>
				</view>
			</view>
			<text v-if="errors.dayMap" class="field-error">{{ errors.dayMap }}</text>

			<text class="card-label" style="margin-top: 24rpx;">高碳日目标</text>
			<view class="macro-row-compact">
				<view class="mrc-item"><text class="mrc-label">热量</text><input type="digit" v-model="form.highCalories" class="mrc-input" /><text class="mrc-unit">kcal</text></view>
				<view class="mrc-item"><text class="mrc-label">蛋白</text><input type="digit" v-model="form.highProtein" class="mrc-input" /><text class="mrc-unit">g</text></view>
				<view class="mrc-item"><text class="mrc-label">碳水</text><input type="digit" v-model="form.highCarbs" class="mrc-input" /><text class="mrc-unit">g</text></view>
				<view class="mrc-item"><text class="mrc-label">脂肪</text><input type="digit" v-model="form.highFat" class="mrc-input" /><text class="mrc-unit">g</text></view>
			</view>

			<text class="card-label" style="margin-top: 24rpx;">低碳日目标</text>
			<view class="macro-row-compact">
				<view class="mrc-item"><text class="mrc-label">热量</text><input type="digit" v-model="form.lowCalories" class="mrc-input" /><text class="mrc-unit">kcal</text></view>
				<view class="mrc-item"><text class="mrc-label">蛋白</text><input type="digit" v-model="form.lowProtein" class="mrc-input" /><text class="mrc-unit">g</text></view>
				<view class="mrc-item"><text class="mrc-label">碳水</text><input type="digit" v-model="form.lowCarbs" class="mrc-input" /><text class="mrc-unit">g</text></view>
				<view class="mrc-item"><text class="mrc-label">脂肪</text><input type="digit" v-model="form.lowFat" class="mrc-input" /><text class="mrc-unit">g</text></view>
			</view>
		</view>

		<!-- 碳水渐降配置 -->
		<view class="card" v-if="form.type === 'carb_taper'">
			<text class="card-label">碳水渐降配置</text>
			<view class="taper-fields">
				<view class="taper-field">
					<text class="taper-label">总周数</text>
					<view class="taper-input-wrap">
						<input type="number" v-model="form.totalWeeks" class="taper-input" placeholder="8" />
						<text class="taper-suffix">周</text>
					</view>
					<text v-if="errors.totalWeeks" class="field-error">{{ errors.totalWeeks }}</text>
				</view>
				<view class="taper-field">
					<text class="taper-label">每周递减</text>
					<view class="taper-input-wrap">
						<input type="digit" v-model="form.weeklyReduction" class="taper-input" placeholder="15" />
						<text class="taper-suffix">g碳水</text>
					</view>
					<text v-if="errors.weeklyReduction" class="field-error">{{ errors.weeklyReduction }}</text>
				</view>
			</view>

			<text class="card-label" style="margin-top: 24rpx;">初始每日目标</text>
			<view class="macro-grid">
				<view class="macro-field">
					<text class="macro-label">热量</text>
					<view class="macro-input-wrap"><input type="digit" v-model="form.calories" class="macro-input" placeholder="0" /><text class="macro-suffix">kcal</text></view>
				</view>
				<view class="macro-field">
					<text class="macro-label">蛋白质</text>
					<view class="macro-input-wrap"><input type="digit" v-model="form.protein" class="macro-input" placeholder="0" /><text class="macro-suffix">g</text></view>
				</view>
				<view class="macro-field">
					<text class="macro-label">碳水</text>
					<view class="macro-input-wrap"><input type="digit" v-model="form.carbs" class="macro-input" placeholder="0" /><text class="macro-suffix">g</text></view>
				</view>
				<view class="macro-field">
					<text class="macro-label">脂肪</text>
					<view class="macro-input-wrap"><input type="digit" v-model="form.fat" class="macro-input" placeholder="0" /><text class="macro-suffix">g</text></view>
				</view>
			</view>

			<!-- 周预览 -->
			<view class="taper-preview" v-if="taperPreview.length > 0">
				<text class="card-label" style="margin-top: 24rpx;">碳水递减预览</text>
				<view v-for="(w, i) in taperPreview" :key="i" class="taper-week">
					<text class="tw-label">第{{ w.week }}周</text>
					<view class="tw-bar-track"><view class="tw-bar-fill" :style="{ width: w.pct + '%' }"></view></view>
					<text class="tw-val" :class="{ 'tw-min': w.capped }">{{ w.carbs }}g</text>
				</view>
			</view>
		</view>

		</template>

		<!-- 保存按钮 -->
		<view class="bottom-action" v-if="templateId || mode === 'custom' || mode === 'edit'">
			<view class="save-btn" :class="{ loading: saving }" @tap="submit">
				<text>{{ saving ? '保存中...' : '保存方案' }}</text>
			</view>
		</view>
	</view>
</template>

<script>
	import { useDietPlanStore } from '../../store/dietPlan'
	import { useUserStore } from '../../store'
	import { PUBLIC_TEMPLATES } from '../../utils/dietPlanTemplates'
	import { calculateMacrosFromTemplate, validatePlanInput, validateCycleConfig, validateTaperConfig } from '../../utils/dietPlanCalc'

	export default {
		data() {
			return {
				mode: 'custom', templateId: '', planId: '',
				selectedPreset: '',
				saving: false,
				errors: {},
				weekDays: ['日', '一', '二', '三', '四', '五', '六'],
				systemTemplates: PUBLIC_TEMPLATES,
				form: {
					name: '', type: 'fixed',
					calories: '', protein: '', carbs: '', fat: '',
					// carb_cycle
					dayMap: { 0: 'low', 1: 'high', 2: 'low', 3: 'high', 4: 'low', 5: 'high', 6: 'low' },
					highCalories: '', highProtein: '', highCarbs: '', highFat: '',
					lowCalories: '', lowProtein: '', lowCarbs: '', lowFat: '',
					// carb_taper
					totalWeeks: '8', weeklyReduction: '15'
				}
			}
		},
		computed: {
			topPadding() {
				const app = getApp()
				return (app.globalData?.customBarHeight || 88) + 8
			},
			pageTitle() {
				if (this.mode === 'edit') return '编辑方案'
				if (this.mode === 'system' && !this.templateId) return '选择方案类型'
				if (this.mode === 'system' || this.mode === 'template') return '创建方案'
				return '自定义方案'
			},
			userWeight() {
				const userStore = useUserStore()
				return userStore.userInfo?.weight || 70
			},
			taperPreview() {
				const total = parseInt(this.form.totalWeeks) || 0
				const reduction = parseFloat(this.form.weeklyReduction) || 0
				const initial = parseFloat(this.form.carbs) || 0
				if (total <= 0 || initial <= 0) return []
				const weeks = []
				for (let i = 1; i <= Math.min(total, 12); i++) {
					const carbs = Math.max(initial - (i - 1) * reduction, 50)
					weeks.push({ week: i, carbs: Math.round(carbs), pct: Math.round((carbs / initial) * 100), capped: carbs <= 50 })
				}
				return weeks
			}
		},
		onLoad(opts) {
			this.mode = opts.mode || 'custom'
			this.templateId = opts.templateId || ''
			this.planId = opts.planId || ''

			const store = useDietPlanStore()
			store.init()

			if (this.mode === 'template' && this.templateId) {
				this.loadTemplate()
			} else if (this.mode === 'edit' && this.planId) {
				this.loadPlan()
			}
		},
		methods: {
			goBack() { uni.navigateBack() },

			typeLabel(type) {
				return { fixed: '固定比例', carb_cycle: '碳循环', carb_taper: '碳水渐降' }[type] || type
			},

			selectPreset(tpl) {
				this.selectedPreset = tpl.id
				this.templateId = tpl.id
				this.loadTemplate()
			},

			loadTemplate() {
				const store = useDietPlanStore()
				const tpl = store.allTemplates.find(t => t.id === this.templateId)
				if (!tpl) return
				this.form.name = tpl.name
				this.form.type = tpl.type

				if (tpl.type === 'carb_cycle' && tpl.defaultCycleConfig) {
					this.form.dayMap = { ...tpl.defaultCycleConfig }
				}
				if (tpl.type === 'carb_taper' && tpl.defaultTaperConfig) {
					this.form.totalWeeks = String(tpl.defaultTaperConfig.totalWeeks)
					this.form.weeklyReduction = String(tpl.defaultTaperConfig.weeklyReduction)
				}

				// 直接用用户体重计算
				this.recalculate()
			},

			loadPlan() {
				const store = useDietPlanStore()
				const plan = store.plans.find(p => p.id === this.planId)
				if (!plan) return
				this.form.name = plan.name
				this.form.type = plan.type
				this.form.calories = String(plan.macroTargets.calories)
				this.form.protein = String(plan.macroTargets.protein)
				this.form.carbs = String(plan.macroTargets.carbs)
				this.form.fat = String(plan.macroTargets.fat)

				if (plan.cycleConfig) {
					this.form.dayMap = { ...plan.cycleConfig.dayMap }
					const h = plan.cycleConfig.highDayTargets || {}
					const l = plan.cycleConfig.lowDayTargets || {}
					this.form.highCalories = String(h.calories || '')
					this.form.highProtein = String(h.protein || '')
					this.form.highCarbs = String(h.carbs || '')
					this.form.highFat = String(h.fat || '')
					this.form.lowCalories = String(l.calories || '')
					this.form.lowProtein = String(l.protein || '')
					this.form.lowCarbs = String(l.carbs || '')
					this.form.lowFat = String(l.fat || '')
				}
				if (plan.taperConfig) {
					this.form.totalWeeks = String(plan.taperConfig.totalWeeks)
					this.form.weeklyReduction = String(plan.taperConfig.weeklyReduction)
				}
			},

			recalculate() {
				const store = useDietPlanStore()
				const tpl = store.allTemplates.find(t => t.id === this.templateId)
				if (!tpl) return
				const w = this.userWeight
				const macros = calculateMacrosFromTemplate(tpl, w)

				if (tpl.type === 'carb_cycle') {
					this.form.highCalories = String(macros.high?.calories || '')
					this.form.highProtein = String(macros.high?.protein || '')
					this.form.highCarbs = String(macros.high?.carbs || '')
					this.form.highFat = String(macros.high?.fat || '')
					this.form.lowCalories = String(macros.low?.calories || '')
					this.form.lowProtein = String(macros.low?.protein || '')
					this.form.lowCarbs = String(macros.low?.carbs || '')
					this.form.lowFat = String(macros.low?.fat || '')
					// 用高碳日作为基础
					this.form.calories = String(macros.high?.calories || '')
					this.form.protein = String(macros.high?.protein || '')
					this.form.carbs = String(macros.high?.carbs || '')
					this.form.fat = String(macros.high?.fat || '')
				} else {
					this.form.calories = String(macros.calories || '')
					this.form.protein = String(macros.protein || '')
					this.form.carbs = String(macros.carbs || '')
					this.form.fat = String(macros.fat || '')
				}
			},

			toggleDay(i) {
				this.form.dayMap[i] = this.form.dayMap[i] === 'high' ? 'low' : 'high'
			},

			validate() {
				this.errors = {}
				const type = this.form.type

				if (!this.form.name.trim()) {
					this.errors.name = '请输入方案名称'
				}

				if (type === 'carb_cycle') {
					const cv = validateCycleConfig(this.form.dayMap)
					if (!cv.valid) Object.assign(this.errors, cv.errors)
				} else if (type === 'carb_taper') {
					const tv = validateTaperConfig({ totalWeeks: parseInt(this.form.totalWeeks), weeklyReduction: parseFloat(this.form.weeklyReduction) })
					if (!tv.valid) Object.assign(this.errors, tv.errors)
				}

				if (type !== 'carb_cycle') {
					const pv = validatePlanInput({ calories: parseFloat(this.form.calories), protein: parseFloat(this.form.protein), carbs: parseFloat(this.form.carbs), fat: parseFloat(this.form.fat) })
					if (!pv.valid) Object.assign(this.errors, pv.errors)
				}

				return Object.keys(this.errors).length === 0
			},

			submit() {
				if (!this.validate()) return
				this.saving = true
				const store = useDietPlanStore()

				try {
					if (this.mode === 'edit') {
						const updates = { name: this.form.name.trim() }
						if (this.form.type === 'carb_cycle') {
							updates.macroTargets = { calories: parseFloat(this.form.highCalories), protein: parseFloat(this.form.highProtein), carbs: parseFloat(this.form.highCarbs), fat: parseFloat(this.form.highFat) }
							updates.cycleConfig = {
								dayMap: { ...this.form.dayMap },
								highDayTargets: { calories: parseFloat(this.form.highCalories), protein: parseFloat(this.form.highProtein), carbs: parseFloat(this.form.highCarbs), fat: parseFloat(this.form.highFat) },
								lowDayTargets: { calories: parseFloat(this.form.lowCalories), protein: parseFloat(this.form.lowProtein), carbs: parseFloat(this.form.lowCarbs), fat: parseFloat(this.form.lowFat) }
							}
						} else if (this.form.type === 'carb_taper') {
							updates.macroTargets = { calories: parseFloat(this.form.calories), protein: parseFloat(this.form.protein), carbs: parseFloat(this.form.carbs), fat: parseFloat(this.form.fat) }
							updates.taperConfig = { totalWeeks: parseInt(this.form.totalWeeks), weeklyReduction: parseFloat(this.form.weeklyReduction), initialCarbs: parseFloat(this.form.carbs), minCarbs: 50 }
						} else {
							updates.macroTargets = { calories: parseFloat(this.form.calories), protein: parseFloat(this.form.protein), carbs: parseFloat(this.form.carbs), fat: parseFloat(this.form.fat) }
						}
						store.updatePlan(this.planId, updates)
					} else if (this.mode === 'template' || this.mode === 'system') {
						const overrides = { name: this.form.name.trim() }
						if (this.form.type === 'carb_cycle') {
							overrides.dayMap = { ...this.form.dayMap }
							overrides.highDayTargets = { calories: parseFloat(this.form.highCalories), protein: parseFloat(this.form.highProtein), carbs: parseFloat(this.form.highCarbs), fat: parseFloat(this.form.highFat) }
							overrides.lowDayTargets = { calories: parseFloat(this.form.lowCalories), protein: parseFloat(this.form.lowProtein), carbs: parseFloat(this.form.lowCarbs), fat: parseFloat(this.form.lowFat) }
						} else if (this.form.type === 'carb_taper') {
							overrides.macroTargets = { calories: parseFloat(this.form.calories), protein: parseFloat(this.form.protein), carbs: parseFloat(this.form.carbs), fat: parseFloat(this.form.fat) }
							overrides.totalWeeks = parseInt(this.form.totalWeeks)
							overrides.weeklyReduction = parseFloat(this.form.weeklyReduction)
							overrides.initialCarbs = parseFloat(this.form.carbs)
						} else {
							overrides.macroTargets = { calories: parseFloat(this.form.calories), protein: parseFloat(this.form.protein), carbs: parseFloat(this.form.carbs), fat: parseFloat(this.form.fat) }
						}
						const plan = store.createPlanFromTemplate(this.templateId, overrides, this.userWeight)
						store.activatePlan(plan.id)
					} else {
						const plan = store.createCustomPlan({ name: this.form.name.trim(), calories: parseFloat(this.form.calories), protein: parseFloat(this.form.protein), carbs: parseFloat(this.form.carbs), fat: parseFloat(this.form.fat) })
						store.activatePlan(plan.id)
					}
					uni.showToast({ title: '保存成功', icon: 'success' })
					setTimeout(() => uni.navigateBack(), 800)
				} catch (e) {
					uni.showToast({ title: e.message || '保存失败', icon: 'none' })
				} finally { this.saving = false }
			}
		}
	}
</script>

<style lang="scss" scoped>
@import '../../styles/variables.scss';

.page { padding: 0 $spacing-xl; padding-bottom: 160rpx; min-height: 100vh; background: $color-bg; }

.nav-bar { @include nav-bar;
	.nav-back { font-size: $font-callout; color: $color-primary; font-weight: 500; }
	.nav-title { font-size: $font-headline; font-weight: 600; color: $color-label; }
	.nav-placeholder { width: 80rpx; }
}

.card { @include card; margin-bottom: $spacing-md; .card-label { @include card-label; } }

// Preset selector
.preset-card {
	padding: $spacing-lg;
	background: $color-fill;
	border-radius: $radius-lg;
	margin-bottom: $spacing-sm;
	border: 2rpx solid transparent;
	transition: all 0.25s cubic-bezier(0.16, 1, 0.3, 1);
	@include press-effect;

	&.selected {
		background: $color-primary-light;
		border-color: $color-primary;
	}

	.preset-top {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: $spacing-xs;
	}
	.preset-name { font-size: $font-body; font-weight: 600; color: $color-label; }
	.preset-type { font-size: $font-caption2; color: $color-label-quaternary; background: $color-bg-elevated; padding: 4rpx $spacing-sm; border-radius: $spacing-xs; }
	.preset-desc { font-size: $font-caption1; color: $color-label-quaternary; line-height: 1.5; }
}

.input-row { @include input-field; .field-input { flex: 1; font-size: $font-body; color: $color-label; } .input-suffix { font-size: $font-subhead; color: $color-label-quaternary; } }

.field-error { display: block; font-size: $font-caption2; color: $color-red; margin-top: $spacing-xs; }

// Macro grid
.macro-grid { display: flex; flex-wrap: wrap; gap: $spacing-sm; }
.macro-field {
	width: calc(50% - #{$spacing-sm} / 2);
	.macro-label { display: block; font-size: $font-caption2; color: $color-label-quaternary; font-weight: 500; margin-bottom: $spacing-xs; }
}
.macro-input-wrap {
	display: flex; align-items: center; background: $color-fill; border-radius: $radius-md; padding: $spacing-md;
	border: 1.5rpx solid transparent; transition: all 0.2s ease;
	&:focus-within { background: rgba(16, 185, 129, 0.04); border-color: rgba(16, 185, 129, 0.3); }
	.macro-input { flex: 1; font-size: $font-headline; font-weight: 700; color: $color-label; text-align: center; font-variant-numeric: tabular-nums; }
	.macro-suffix { font-size: $font-caption1; color: $color-label-quaternary; margin-left: 4rpx; }
}

// Cycle config
.cycle-week { display: flex; gap: $spacing-xs; }
.cycle-day {
	flex: 1; text-align: center; padding: $spacing-md $spacing-xs; border-radius: $radius-md;
	background: $color-fill; transition: all 0.2s ease; @include press-effect;
	.cycle-day-name { display: block; font-size: $font-caption1; color: $color-label-tertiary; font-weight: 500; }
	.cycle-day-type { display: block; font-size: $font-caption2; color: $color-label-quaternary; margin-top: 4rpx; }
	&.high { background: $color-primary-light; .cycle-day-name { color: $color-primary; } .cycle-day-type { color: $color-primary; } }
	&.low { background: $color-fill; }
}

// Compact macro row
.macro-row-compact { display: flex; gap: $spacing-xs; }
.mrc-item {
	flex: 1; text-align: center;
	.mrc-label { display: block; font-size: $font-caption2; color: $color-label-quaternary; margin-bottom: 4rpx; }
	.mrc-input { width: 100%; background: $color-fill; border-radius: $radius-sm; padding: $spacing-sm; text-align: center; font-size: $font-subhead; font-weight: 600; color: $color-label; font-variant-numeric: tabular-nums; }
	.mrc-unit { display: block; font-size: $font-caption2; color: $color-label-quaternary; margin-top: 2rpx; }
}

// Taper fields
.taper-fields { display: flex; gap: $spacing-md; }
.taper-field {
	flex: 1;
	.taper-label { display: block; font-size: $font-caption2; color: $color-label-quaternary; font-weight: 500; margin-bottom: $spacing-xs; }
}
.taper-input-wrap {
	display: flex; align-items: center; background: $color-fill; border-radius: $radius-md; padding: $spacing-md;
	.taper-input { flex: 1; font-size: $font-headline; font-weight: 700; color: $color-label; text-align: center; font-variant-numeric: tabular-nums; }
	.taper-suffix { font-size: $font-caption1; color: $color-label-quaternary; }
}

// Taper preview
.taper-week {
	display: flex; align-items: center; margin-bottom: $spacing-sm;
	.tw-label { width: 100rpx; font-size: $font-caption1; color: $color-label-quaternary; font-weight: 500; }
	.tw-bar-track { flex: 1; height: 16rpx; background: $color-fill; border-radius: 8rpx; margin: 0 $spacing-sm; overflow: hidden; }
	.tw-bar-fill { height: 100%; background: linear-gradient(90deg, $color-primary, $color-mint); border-radius: 8rpx; transition: width 0.3s ease; }
	.tw-val { width: 80rpx; text-align: right; font-size: $font-caption1; font-weight: 600; color: $color-label; font-variant-numeric: tabular-nums; }
	.tw-min { color: $color-red; }
}

.bottom-action { @include bottom-action-bar; .save-btn { @include primary-button; } }
</style>
