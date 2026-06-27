import React from 'react'
import { Layout as AntLayout } from 'antd'
import Header from './Header'
import Footer from './Footer'

interface LayoutProps {
  children: React.ReactNode
}

const Layout: React.FC<LayoutProps> = ({ children }) => {
  return (
    <AntLayout style={{ minHeight: '100vh' }}>
      <Header />
      <main>{children}</main>
      <Footer />
    </AntLayout>
  )
}

export default Layout
