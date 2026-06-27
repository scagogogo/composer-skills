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
    <div style={{ textAlign: 'center', marginBottom: 56 }}>
      {id && <div id={id} style={{ position: 'relative', top: -100 }} />}
      <Title
        level={2}
        style={{
          marginBottom: 14,
          fontWeight: 800,
          fontSize: 'clamp(24px, 3vw, 36px)',
          letterSpacing: '-0.01em',
          color: '#1E293B',
        }}
      >
        {title}
      </Title>
      {subtitle && (
        <Paragraph
          style={{
            fontSize: 16,
            color: '#64748B',
            maxWidth: 600,
            margin: '0 auto',
            lineHeight: 1.6,
          }}
        >
          {subtitle}
        </Paragraph>
      )}
      <div
        style={{
          width: 48,
          height: 4,
          borderRadius: 2,
          background: 'linear-gradient(135deg, #4F46E5, #0284C7)',
          margin: '20px auto 0',
        }}
      />
    </div>
  )
}

export default SectionTitle
