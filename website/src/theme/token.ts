import type { ThemeConfig } from 'antd'

export const themeConfig: ThemeConfig = {
  token: {
    colorPrimary: '#4F46E5',
    colorInfo: '#0284C7',
    colorSuccess: '#059669',
    colorWarning: '#D97706',
    colorError: '#E11D48',
    colorTextBase: '#1E293B',
    colorTextSecondary: '#475569',
    borderRadius: 8,
    fontFamily: "'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif",
  },
  components: {
    Button: {
      primaryShadow: '0 4px 14px 0 rgba(79, 70, 229, 0.35)',
    },
    Card: {
      borderRadiusLG: 12,
    },
  },
}
