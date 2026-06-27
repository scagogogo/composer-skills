import React from 'react'
import { useTranslation } from 'react-i18next'
import { Row, Col, Tag } from 'antd'
import { CloseCircleOutlined, CheckCircleOutlined, ArrowRightOutlined } from '@ant-design/icons'
import SectionTitle from '../components/SectionTitle'
import CodeBlock from '../components/CodeBlock'
import { RevealSection, StaggerGrid, StaggerItem } from '../components/ScrollReveal'

const ProblemSection: React.FC = () => {
  const { t } = useTranslation()

  const painPoints = t('problem.painPoints', { returnObjects: true }) as Array<{
    pain: string
    solution: string
  }>

  return (
    <section id="problem" className="section-alt">
      <div style={{ maxWidth: 1100, margin: '0 auto' }}>
        <SectionTitle title={t('problem.title')} subtitle={t('problem.subtitle')} />

        {/* Code comparison */}
        <RevealSection style={{ marginBottom: 48 }}>
          <Row gutter={[20, 20]} align="stretch">
            <Col xs={24} md={12}>
              <div
                style={{
                  border: '1px solid #FECACA',
                  background: '#FEF2F2',
                  borderRadius: 4,
                  padding: '16px 20px 12px',
                  height: '100%',
                }}
              >
                <div style={{ display: 'flex', alignItems: 'center', gap: 8, marginBottom: 10 }}>
                  <Tag color="error" style={{ margin: 0, borderRadius: 2, fontWeight: 600 }}>OLD</Tag>
                  <span style={{ color: '#991B1B', fontWeight: 600, fontSize: 13 }}>{t('problem.oldWay')}</span>
                </div>
                <div className="code-wrapper">
                  <CodeBlock code={t('problem.oldCode')} />
                </div>
              </div>
            </Col>
            <Col xs={24} md={12}>
              <div
                style={{
                  border: '1px solid #BBF7D0',
                  background: '#F0FDF4',
                  borderRadius: 4,
                  padding: '16px 20px 12px',
                  height: '100%',
                }}
              >
                <div style={{ display: 'flex', alignItems: 'center', gap: 8, marginBottom: 10 }}>
                  <Tag color="success" style={{ margin: 0, borderRadius: 2, fontWeight: 600 }}>NEW</Tag>
                  <span style={{ color: '#166534', fontWeight: 600, fontSize: 13 }}>{t('problem.newWay')}</span>
                </div>
                <div className="code-wrapper">
                  <CodeBlock code={t('problem.newCode')} />
                </div>
              </div>
            </Col>
          </Row>
        </RevealSection>

        {/* Pain → Solution cards */}
        <StaggerGrid>
          <div style={{ display: 'grid', gap: 12 }}>
            {painPoints.map((item, index) => (
              <StaggerItem key={index}>
                <div
                  style={{
                    display: 'flex',
                    alignItems: 'center',
                    gap: 12,
                    padding: '14px 18px',
                    background: '#fff',
                    border: '1px solid #E2E8F0',
                    borderRadius: 4,
                    flexWrap: 'wrap',
                  }}
                >
                  <div style={{ display: 'flex', alignItems: 'center', gap: 8, flex: '1 1 200px', minWidth: 200 }}>
                    <CloseCircleOutlined style={{ color: '#DC2626', fontSize: 16, flexShrink: 0 }} />
                    <span style={{ fontSize: 14, color: '#64748B' }}>{item.pain}</span>
                  </div>
                  <ArrowRightOutlined style={{ color: '#CBD5E1', fontSize: 12, flexShrink: 0 }} />
                  <div style={{ display: 'flex', alignItems: 'center', gap: 8, flex: '1 1 200px', minWidth: 200 }}>
                    <CheckCircleOutlined style={{ color: '#16A34A', fontSize: 16, flexShrink: 0 }} />
                    <span style={{ fontSize: 14, color: '#0F172A', fontWeight: 500 }}>{item.solution}</span>
                  </div>
                </div>
              </StaggerItem>
            ))}
          </div>
        </StaggerGrid>
      </div>
    </section>
  )
}

export default ProblemSection
