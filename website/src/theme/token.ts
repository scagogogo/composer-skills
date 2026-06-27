import type { ThemeConfig } from 'antd'

export const themeConfig: ThemeConfig = {
  token: {
    colorPrimary: '#2563EB',
    colorInfo: '#0891B2',
    colorSuccess: '#16A34A',
    colorWarning: '#EA580C',
    colorError: '#DC2626',
    colorTextBase: '#1E293B',
    colorTextSecondary: '#64748B',
    colorBgContainer: '#FFFFFF',
    colorBgLayout: '#FFFFFF',
    borderRadius: 4,
    fontFamily: "'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif",
    fontSize: 15,
    lineHeight: 1.6,
  },
  components: {
    Button: {
      primaryShadow: 'none',
      borderRadius: 4,
      controlHeight: 38,
      fontWeight: 600,
    },
    Card: {
      borderRadiusLG: 4,
      paddingLG: 20,
    },
    Table: {
      borderRadius: 0,
    },
    Tabs: {
      inkBarColor: '#2563EB',
      itemActiveColor: '#2563EB',
      itemSelectedColor: '#2563EB',
      itemHoverColor: '#1D4ED8',
    },
    Menu: {
      itemSelectedColor: '#2563EB',
      itemHoverColor: '#1D4ED8',
    },
    Tag: {
      borderRadiusSM: 2,
    },
  },
}
