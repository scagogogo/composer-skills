import React from 'react'
import { Typography } from 'antd'

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
      <Title
        level={2}
        style={{
          marginBottom: 10,
          fontWeight: 800,
          fontSize: 'clamp(22px, 2.8vw, 32px)',
          color: '#0F172A',
        }}
      >
        {title}
      </Title>
      {subtitle && (
        <Paragraph
          style={{
            fontSize: 15,
            color: '#64748B',
            maxWidth: 560,
            margin: '0 auto',
            lineHeight: 1.6,
          }}
        >
          {subtitle}
        </Paragraph>
      )}
      <div
        style={{
          width: 32,
          height: 3,
          background: '#2563EB',
          margin: '16px auto 0',
        }}
      />
    </div>
  )
}

export default SectionTitle
