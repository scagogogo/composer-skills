import React from 'react'
import { useTranslation } from 'react-i18next'
import { Layout as AntLayout, Menu, Button, Drawer } from 'antd'
import { GithubOutlined, MenuOutlined } from '@ant-design/icons'
import LangSwitch from '../components/LangSwitch'

const { Header: AntHeader } = AntLayout

const navItems = ['features', 'architecture', 'quickStart', 'security', 'coverage'] as const

const Header: React.FC = () => {
  const { t } = useTranslation()
  const [drawerOpen, setDrawerOpen] = React.useState(false)

  const menuItems = navItems.map((key) => ({
    key,
    label: (
      <a
        href={`#${key}`}
        style={{
          fontWeight: 500,
          fontSize: 14,
          color: '#475569',
        }}
      >
        {t(`nav.${key}`)}
      </a>
    ),
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
        borderBottom: '1px solid #E2E8F0',
        height: 56,
      }}
    >
      <div style={{ display: 'flex', alignItems: 'center', gap: 8 }}>
        <div
          style={{
            width: 28,
            height: 28,
            borderRadius: 4,
            background: '#2563EB',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            color: '#fff',
            fontWeight: 800,
            fontSize: 13,
            fontFamily: 'monospace',
          }}
        >
          CS
        </div>
        <span style={{ fontWeight: 700, fontSize: 16, color: '#0F172A' }}>
          Composer Skills
        </span>
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

      <div style={{ display: 'flex', alignItems: 'center', gap: 10 }}>
        <LangSwitch />
        <Button
          type="text"
          icon={<GithubOutlined />}
          href="https://github.com/scagogogo/composer-skills"
          target="_blank"
          style={{ fontSize: 18, color: '#475569' }}
        />
        <Button
          className="mobile-menu-btn"
          type="text"
          icon={<MenuOutlined />}
          onClick={() => setDrawerOpen(true)}
          style={{ color: '#1E293B' }}
        />
      </div>

      <Drawer
        title={
          <div style={{ display: 'flex', alignItems: 'center', gap: 8 }}>
            <div
              style={{
                width: 24,
                height: 24,
                borderRadius: 3,
                background: '#2563EB',
                display: 'flex',
                alignItems: 'center',
                justifyContent: 'center',
                color: '#fff',
                fontWeight: 800,
                fontSize: 11,
                fontFamily: 'monospace',
              }}
            >
              CS
            </div>
            <span style={{ fontWeight: 700, fontSize: 15 }}>Composer Skills</span>
          </div>
        }
        placement="right"
        onClose={() => setDrawerOpen(false)}
        open={drawerOpen}
        width={260}
      >
        <Menu
          mode="vertical"
          items={navItems.map((key) => ({
            key,
            label: (
              <a href={`#${key}`} onClick={handleMenuClick} style={{ fontWeight: 500 }}>
                {t(`nav.${key}`)}
              </a>
            ),
          }))}
        />
        <div style={{ marginTop: 12, padding: '0 16px' }}>
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
