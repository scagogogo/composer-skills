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
          transition: 'color 0.2s',
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
        padding: '0 32px',
        background: 'rgba(255, 255, 255, 0.92)',
        backdropFilter: 'blur(12px)',
        borderBottom: '1px solid rgba(0, 0, 0, 0.06)',
        height: 64,
      }}
    >
      <div style={{ display: 'flex', alignItems: 'center', gap: 10 }}>
        <div
          style={{
            width: 32,
            height: 32,
            borderRadius: 8,
            background: 'linear-gradient(135deg, #4F46E5, #6366F1)',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            color: '#fff',
            fontWeight: 800,
            fontSize: 14,
            fontFamily: 'monospace',
          }}
        >
          CS
        </div>
        <span
          style={{
            fontWeight: 700,
            fontSize: 18,
            background: 'linear-gradient(135deg, #1E293B, #4F46E5)',
            WebkitBackgroundClip: 'text',
            WebkitTextFillColor: 'transparent',
          }}
        >
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

      <div style={{ display: 'flex', alignItems: 'center', gap: 12 }}>
        <LangSwitch />
        <Button
          type="link"
          icon={<GithubOutlined />}
          href="https://github.com/scagogogo/composer-skills"
          target="_blank"
          style={{ fontSize: 20, color: '#475569' }}
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
                width: 28,
                height: 28,
                borderRadius: 6,
                background: 'linear-gradient(135deg, #4F46E5, #6366F1)',
                display: 'flex',
                alignItems: 'center',
                justifyContent: 'center',
                color: '#fff',
                fontWeight: 800,
                fontSize: 12,
                fontFamily: 'monospace',
              }}
            >
              CS
            </div>
            <span style={{ fontWeight: 700 }}>Composer Skills</span>
          </div>
        }
        placement="right"
        onClose={() => setDrawerOpen(false)}
        open={drawerOpen}
        width={280}
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
