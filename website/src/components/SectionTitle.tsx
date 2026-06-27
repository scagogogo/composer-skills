import React from 'react'
import { Typography, Divider } from 'antd'

const { Title, Paragraph } = Typography

interface SectionTitleProps {
  id?: string
  title: string
  subtitle?: string
}

const SectionTitle: React.FC<SectionTitleProps> = ({ id, title, subtitle }) => {
  return (
    <div style={{ textAlign: 'center', marginBottom: 48 }}>
      {id && <div id={id} style={{ position: 'relative', top: -80 }} />}
      <Title level={2} style={{ marginBottom: 12 }}>
        {title}
      </Title>
      {subtitle && (
        <Paragraph style={{ fontSize: 16, color: '#475569', maxWidth: 640, margin: '0 auto' }}>
          {subtitle}
        </Paragraph>
      )}
      <Divider
        style={{
          width: 60,
          minWidth: 60,
          margin: '16px auto 0',
          borderColor: '#4F46E5',
          borderWidth: 2,
        }}
      />
    </div>
  )
}

export default SectionTitle
