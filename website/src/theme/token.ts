import type { ThemeConfig } from 'antd'

export const themeConfig: ThemeConfig = {
  token: {
    colorPrimary: '#4F46E5',
    colorInfo: '#0284C7',
    colorSuccess: '#059669',
    colorWarning: '#D97706',
    colorError: '#E11D48',
    colorTextBase: '#1E293B',
    colorTextSecondary: '#64748B',
    colorBgContainer: '#FFFFFF',
    colorBgLayout: '#FFFFFF',
    borderRadius: 10,
    fontFamily: "'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif",
    fontSize: 15,
    lineHeight: 1.6,
  },
  components: {
    Button: {
      primaryShadow: '0 4px 14px 0 rgba(79, 70, 229, 0.35)',
      borderRadius: 8,
      controlHeight: 40,
      fontWeight: 600,
    },
    Card: {
      borderRadiusLG: 16,
      paddingLG: 24,
    },
    Table: {
      borderRadius: 12,
    },
    Tabs: {
      inkBarColor: '#4F46E5',
      itemActiveColor: '#4F46E5',
      itemSelectedColor: '#4F46E5',
      itemHoverColor: '#6366F1',
    },
    Menu: {
      itemSelectedColor: '#4F46E5',
      itemHoverColor: '#6366F1',
    },
  },
}
