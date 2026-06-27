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
      style={{
        height: '100%',
        background: '#FFFFFF',
      }}
      styles={{
        body: {
          padding: 20,
        },
      }}
    >
      <div
        style={{
          width: 40,
          height: 40,
          borderRadius: 4,
          background: '#EFF6FF',
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
          fontSize: 20,
          color: '#2563EB',
          marginBottom: 16,
        }}
      >
        {icon}
      </div>
      <Title level={5} style={{ marginBottom: 8, fontWeight: 700, fontSize: 15 }}>
        {title}
      </Title>
      <Paragraph style={{ color: '#64748B', marginBottom: 0, lineHeight: 1.6, fontSize: 14 }}>
        {description}
      </Paragraph>
    </Card>
  )
}

export default FeatureCard
