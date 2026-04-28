/**
 * Global test setup — provides uni-app API mock for all tests
 */
const mockStorage = {}

globalThis.uni = {
  setStorageSync(key, value) {
    mockStorage[key] = value
  },
  getStorageSync(key) {
    return mockStorage[key] || ''
  },
  removeStorageSync(key) {
    delete mockStorage[key]
  },
  showToast() {},
  showModal() {},
  navigateTo() {},
  navigateBack() {},
  switchTab() {}
}

globalThis.__testMockStorage = mockStorage
