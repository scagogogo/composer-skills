import React from 'react'
import { useTranslation } from 'react-i18next'
import { Layout as AntLayout, Menu, Button, Drawer } from 'antd'
import { GithubOutlined, MenuOutlined, GlobalOutlined } from '@ant-design/icons'
import LangSwitch from '../components/LangSwitch'

const { Header: AntHeader } = AntLayout

const navItems = ['features', 'architecture', 'quickStart', 'security', 'coverage'] as const

const Header: React.FC = () => {
  const { t } = useTranslation()
  const [drawerOpen, setDrawerOpen] = React.useState(false)

  const menuItems = navItems.map((key) => ({
    key,
    label: <a href={`#${key}`}>{t(`nav.${key}`)}</a>,
  }))

  const handleMenuClick = () => {
    setDrawerOpen(false)
  }

  return (
    <AntHeader
      style={{
        position: 'sticky',
        top: 0,
        zIndex: 1000,
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'space-between',
        padding: '0 24px',
        background: '#fff',
        boxShadow: '0 2px 8px rgba(0, 0, 0, 0.06)',
      }}
    >
      <div style={{ display: 'flex', alignItems: 'center', gap: 8 }}>
        <GlobalOutlined style={{ fontSize: 24, color: '#4F46E5' }} />
        <span style={{ fontWeight: 700, fontSize: 18, color: '#1E293B' }}>Composer Skills</span>
      </div>

      {/* Desktop menu */}
      <Menu
        mode="horizontal"
        items={menuItems}
        style={{
          flex: 1,
          justifyContent: 'center',
          border: 'none',
          background: 'transparent',
          display: 'flex',
        }}
        className="desktop-menu"
      />

      <div style={{ display: 'flex', alignItems: 'center', gap: 12 }}>
        <LangSwitch />
        <Button
          type="link"
          icon={<GithubOutlined />}
          href="https://github.com/scagogogo/composer-skills"
          target="_blank"
          style={{ fontSize: 20 }}
        />
        <Button
          className="mobile-menu-btn"
          type="text"
          icon={<MenuOutlined />}
          onClick={() => setDrawerOpen(true)}
        />
      </div>

      <Drawer
        title="Composer Skills"
        placement="right"
        onClose={() => setDrawerOpen(false)}
        open={drawerOpen}
        width={280}
      >
        <Menu
          mode="vertical"
          items={navItems.map((key) => ({
            key,
            label: <a href={`#${key}`} onClick={handleMenuClick}>{t(`nav.${key}`)}</a>,
          }))}
        />
        <div style={{ marginTop: 16, padding: '0 16px' }}>
          <Button
            block
            icon={<GithubOutlined />}
            href="https://github.com/scagogogo/composer-skills"
            target="_blank"
          >
            GitHub
          </Button>
        </div>
      </Drawer>
    </AntHeader>
  )
}

export default Header
