import React from 'react'
import { useTranslation } from 'react-i18next'
import { Row, Col } from 'antd'
import { BugOutlined, SafetyOutlined, FileProtectOutlined } from '@ant-design/icons'
import SectionTitle from '../components/SectionTitle'
import CodeBlock from '../components/CodeBlock'
import { RevealSection, StaggerGrid, StaggerItem } from '../components/ScrollReveal'

const securityIcons = [<BugOutlined />, <SafetyOutlined />, <FileProtectOutlined />]
const securityColors = ['#DC2626', '#0891B2', '#2563EB']

const SecuritySection: React.FC = () => {
  const { t } = useTranslation()

  const cards = [
    { title: t('security.auditTitle'), code: t('security.auditCode') },
    { title: t('security.remoteTitle'), code: t('security.remoteCode') },
    { title: t('security.validateTitle'), code: t('security.validateCode') },
  ]

  return (
    <section id="security" className="section-alt">
      <div style={{ maxWidth: 1100, margin: '0 auto' }}>
        <SectionTitle title={t('security.title')} subtitle={t('security.subtitle')} />

        {/* Security features overview — 3 metric bars */}
        <RevealSection style={{ marginBottom: 36 }}>
          <Row gutter={[16, 12]}>
            {[
              { label: 'Local Audit', pct: 100, color: '#16A34A' },
              { label: 'Remote Advisories', pct: 100, color: '#0891B2' },
              { label: 'Schema Validation', pct: 100, color: '#2563EB' },
            ].map((item, i) => (
              <Col xs={24} md={8} key={i}>
                <div style={{ padding: '4px 0' }}>
                  <div style={{ display: 'flex', justifyContent: 'space-between', marginBottom: 6 }}>
                    <span style={{ fontSize: 13, fontWeight: 600, color: '#334155' }}>{item.label}</span>
                    <span style={{ fontSize: 13, fontWeight: 700, color: item.color }}>✓</span>
                  </div>
                  <div style={{ height: 6, background: '#F1F5F9', borderRadius: 2, overflow: 'hidden' }}>
                    <div
                      style={{
                        height: '100%',
                        width: `${item.pct}%`,
                        background: item.color,
                        borderRadius: 2,
                      }}
                    />
                  </div>
                </div>
              </Col>
            ))}
          </Row>
        </RevealSection>

        {/* Code cards */}
        <StaggerGrid>
          <Row gutter={[20, 20]}>
            {cards.map((card, index) => {
              const color = securityColors[index]
              return (
                <Col xs={24} md={8} key={index}>
                  <StaggerItem>
                    <div
                      style={{
                        height: '100%',
                        border: `1px solid ${color}30`,
                        borderRadius: 4,
                        background: '#fff',
                        overflow: 'hidden',
                      }}
                    >
                      {/* Header bar */}
                      <div
                        style={{
                          padding: '10px 16px',
                          background: `${color}08`,
                          borderBottom: `1px solid ${color}20`,
                          display: 'flex',
                          alignItems: 'center',
                          gap: 8,
                        }}
                      >
                        <span style={{ fontSize: 16, color }}>{securityIcons[index]}</span>
                        <span style={{ fontWeight: 600, fontSize: 14, color }}>{card.title}</span>
                      </div>
                      {/* Code */}
                      <div style={{ padding: '0 12px 12px' }}>
                        <CodeBlock code={card.code} />
                      </div>
                    </div>
                  </StaggerItem>
                </Col>
              )
            })}
          </Row>
        </StaggerGrid>
      </div>
    </section>
  )
}

export default SecuritySection
