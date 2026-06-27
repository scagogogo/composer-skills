import React from 'react'
import { useTranslation } from 'react-i18next'
import { Button, Row, Col, Typography } from 'antd'
import { RocketOutlined, GithubOutlined } from '@ant-design/icons'

const { Title, Paragraph } = Typography

const HeroSection: React.FC = () => {
  const { t } = useTranslation()

  const stats = [
    { value: '234', label: t('hero.statMethods') },
    { value: '20', label: t('hero.statApi') },
    { value: '50+', label: t('hero.statCli') },
    { value: '450+', label: t('hero.statTests') },
  ]

  return (
    <section
      id="hero"
      style={{
        background: '#0F172A',
        color: '#fff',
        padding: '100px 24px 80px',
        textAlign: 'center',
      }}
    >
      <div style={{ maxWidth: 860, margin: '0 auto' }}>
        {/* Badge */}
        <div
          className="badge"
          style={{
            borderColor: '#334155',
            color: '#94A3B8',
            background: '#1E293B',
            marginBottom: 24,
          }}
        >
          {t('hero.badge')}
        </div>

        <Title
          level={1}
          style={{
            color: '#F8FAFC',
            fontSize: 'clamp(28px, 5vw, 44px)',
            marginBottom: 16,
            lineHeight: 1.2,
            fontWeight: 800,
          }}
        >
          {t('hero.tagline')}
        </Title>

        <Paragraph
          style={{
            color: '#94A3B8',
            fontSize: 'clamp(15px, 1.8vw, 18px)',
            maxWidth: 640,
            margin: '0 auto 36px',
            lineHeight: 1.7,
          }}
        >
          {t('hero.subtitle')}
        </Paragraph>

        <div style={{ display: 'flex', gap: 12, justifyContent: 'center', flexWrap: 'wrap' }}>
          <Button
            type="primary"
            size="large"
            icon={<RocketOutlined />}
            href="#quickStart"
            style={{
              fontWeight: 700,
              height: 44,
              paddingInline: 28,
              borderRadius: 4,
              fontSize: 15,
            }}
          >
            {t('hero.cta')}
          </Button>
          <Button
            size="large"
            icon={<GithubOutlined />}
            href="https://github.com/scagogogo/composer-skills"
            target="_blank"
            style={{
              background: 'transparent',
              color: '#CBD5E1',
              border: '1px solid #334155',
              fontWeight: 600,
              height: 44,
              paddingInline: 24,
              borderRadius: 4,
              fontSize: 15,
            }}
          >
            {t('hero.ctaSecondary')}
          </Button>
        </div>

        <Row
          gutter={[16, 16]}
          style={{ marginTop: 56, maxWidth: 600, marginLeft: 'auto', marginRight: 'auto' }}
        >
          {stats.map((stat, index) => (
            <Col xs={12} sm={6} key={index}>
              <div className="stat-card">
                <div style={{ fontSize: 28, fontWeight: 800, lineHeight: 1.2, color: '#F8FAFC' }}>
                  {stat.value}
                </div>
                <div
                  style={{
                    fontSize: 12,
                    color: '#64748B',
                    marginTop: 4,
                    fontWeight: 500,
                    textTransform: 'uppercase',
                    letterSpacing: '0.06em',
                  }}
                >
                  {stat.label}
                </div>
              </div>
            </Col>
          ))}
        </Row>
      </div>
    </section>
  )
}

export default HeroSection
