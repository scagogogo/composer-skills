import React from 'react'
import { Card, Typography } from 'antd'

const { Title, Paragraph } = Typography

interface FeatureCardProps {
  icon: React.ReactNode
  title: string
  description: string
}

const FeatureCard: React.FC<FeatureCardProps> = ({ icon, title, description }) => {
  return (
    <Card
      className="feature-card"
      hoverable
      style={{
        height: '100%',
        borderRadius: 16,
        border: '1px solid #E2E8F0',
        background: '#FFFFFF',
      }}
      styles={{
        body: {
          padding: '28px',
        },
      }}
    >
      <div
        style={{
          width: 52,
          height: 52,
          borderRadius: 14,
          background: 'linear-gradient(135deg, #EEF2FF, #E0E7FF)',
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
          fontSize: 24,
          color: '#4F46E5',
          marginBottom: 20,
        }}
      >
        {icon}
      </div>
      <Title level={5} style={{ marginBottom: 10, fontWeight: 700, fontSize: 16 }}>
        {title}
      </Title>
      <Paragraph style={{ color: '#64748B', marginBottom: 0, lineHeight: 1.6, fontSize: 14 }}>
        {description}
      </Paragraph>
    </Card>
  )
}

export default FeatureCard
