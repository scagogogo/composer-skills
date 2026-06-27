import React from 'react'
import { useTranslation } from 'react-i18next'
import { Button, Row, Col, Typography } from 'antd'
import { RocketOutlined, GithubOutlined, DownOutlined } from '@ant-design/icons'

const { Title, Paragraph } = Typography

const HeroSection: React.FC = () => {
  const { t } = useTranslation()

  const stats = [
    { value: '234', label: t('hero.statMethods'), color: '#818CF8' },
    { value: '20', label: t('hero.statApi'), color: '#38BDF8' },
    { value: '50+', label: t('hero.statCli'), color: '#34D399' },
    { value: '450+', label: t('hero.statTests'), color: '#FBBF24' },
  ]

  return (
    <section
      id="hero"
      className="hero-bg"
      style={{
        color: '#fff',
        padding: '120px 24px 100px',
        textAlign: 'center',
        position: 'relative',
        overflow: 'hidden',
        minHeight: '90vh',
        display: 'flex',
        alignItems: 'center',
      }}
    >
      {/* Decorative floating dots */}
      <div
        className="float-dot"
        style={{ width: 300, height: 300, background: '#818CF8', top: -50, left: -100 }}
      />
      <div
        className="float-dot-slow"
        style={{ width: 200, height: 200, background: '#38BDF8', top: '20%', right: -60 }}
      />
      <div
        className="float-dot"
        style={{ width: 150, height: 150, background: '#34D399', bottom: '15%', left: '10%', animationDelay: '2s' }}
      />
      <div
        className="float-dot-slow"
        style={{ width: 180, height: 180, background: '#FBBF24', bottom: '5%', right: '15%', animationDelay: '1s' }}
      />

      {/* Grid pattern overlay */}
      <div
        style={{
          position: 'absolute',
          inset: 0,
          backgroundImage:
            'radial-gradient(circle, rgba(255,255,255,0.05) 1px, transparent 1px)',
          backgroundSize: '32px 32px',
          pointerEvents: 'none',
        }}
      />

      <div style={{ maxWidth: 920, margin: '0 auto', position: 'relative', zIndex: 1 }}>
        {/* Badge */}
        <div
          className="badge"
          style={{
            background: 'rgba(255, 255, 255, 0.12)',
            color: '#E0E7FF',
            border: '1px solid rgba(255, 255, 255, 0.15)',
            marginBottom: 24,
            display: 'inline-flex',
          }}
        >
          🚀 {t('hero.badge')}
        </div>

        <Title
          level={1}
          style={{
            color: '#fff',
            fontSize: 'clamp(30px, 5.5vw, 56px)',
            marginBottom: 20,
            lineHeight: 1.15,
            fontWeight: 800,
            letterSpacing: '-0.02em',
          }}
        >
          {t('hero.tagline')}
        </Title>

        <Paragraph
          style={{
            color: 'rgba(255, 255, 255, 0.82)',
            fontSize: 'clamp(16px, 2vw, 20px)',
            maxWidth: 700,
            margin: '0 auto 40px',
            lineHeight: 1.7,
            fontWeight: 400,
          }}
        >
          {t('hero.subtitle')}
        </Paragraph>

        <div style={{ display: 'flex', gap: 16, justifyContent: 'center', flexWrap: 'wrap' }}>
          <Button
            type="primary"
            size="large"
            icon={<RocketOutlined />}
            href="#quickStart"
            style={{
              background: '#fff',
              color: '#4F46E5',
              border: 'none',
              fontWeight: 700,
              height: 52,
              paddingInline: 36,
              borderRadius: 10,
              fontSize: 16,
              boxShadow: '0 4px 20px rgba(0, 0, 0, 0.15)',
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
              background: 'rgba(255, 255, 255, 0.1)',
              color: '#fff',
              border: '1px solid rgba(255, 255, 255, 0.25)',
              fontWeight: 600,
              height: 52,
              paddingInline: 32,
              borderRadius: 10,
              fontSize: 16,
              backdropFilter: 'blur(8px)',
            }}
          >
            {t('hero.ctaSecondary')}
          </Button>
        </div>

        <Row
          gutter={[20, 20]}
          style={{ marginTop: 72, maxWidth: 720, marginLeft: 'auto', marginRight: 'auto' }}
        >
          {stats.map((stat, index) => (
            <Col xs={12} sm={6} key={index}>
              <div className="stat-card">
                <div
                  style={{
                    fontSize: 36,
                    fontWeight: 800,
                    lineHeight: 1.2,
                    background: `linear-gradient(135deg, #fff, ${stat.color})`,
                    WebkitBackgroundClip: 'text',
                    WebkitTextFillColor: 'transparent',
                  }}
                >
                  {stat.value}
                </div>
                <div
                  style={{
                    fontSize: 13,
                    color: 'rgba(255, 255, 255, 0.7)',
                    marginTop: 6,
                    fontWeight: 500,
                    textTransform: 'uppercase',
                    letterSpacing: '0.05em',
                  }}
                >
                  {stat.label}
                </div>
              </div>
            </Col>
          ))}
        </Row>

        {/* Scroll indicator */}
        <div style={{ marginTop: 48 }} className="scroll-indicator">
          <DownOutlined style={{ fontSize: 20, color: 'rgba(255,255,255,0.4)' }} />
        </div>
      </div>
    </section>
  )
}

export default HeroSection
