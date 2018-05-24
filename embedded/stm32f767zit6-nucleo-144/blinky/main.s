.thumb
.syntax unified

.equ STACKINIT,0x20002000

.section .text
	.org 0

isr_vectors:
  .word  0x20002000
  .word  main
  .word  _handler // NMI_Handler
  .word  _handler // HardFault_Handler
  .word  _handler // MemManage_Handler
  .word  _handler // BusFault_Handler
  .word  _handler // UsageFault_Handler
  .word  0
  .word  0
  .word  0
  .word  0
  .word  _handler // SVC_Handler
  .word  _handler // DebugMon_Handler
  .word  0
  .word  _handler // PendSV_Handler
  .word  _handler // SysTick_Handler
  
  /* External Interrupts */
  .word     _handler // WWDG_IRQHandler                   /* Window WatchDog              */
  .word     _handler // PVD_IRQHandler                    /* PVD through EXTI Line detection */
  .word _handler //     TAMP_STAMP_IRQHandler             /* Tamper and TimeStamps through the EXTI line */
  .word _handler //     RTC_WKUP_IRQHandler               /* RTC Wakeup through the EXTI line */
  .word _handler //     FLASH_IRQHandler                  /* FLASH                        */
  .word _handler //     RCC_IRQHandler                    /* RCC                          */
  .word _handler //     EXTI0_IRQHandler                  /* EXTI Line0                   */
  .word _handler //     EXTI1_IRQHandler                  /* EXTI Line1                   */
  .word _handler //     EXTI2_IRQHandler                  /* EXTI Line2                   */
  .word _handler //     EXTI3_IRQHandler                  /* EXTI Line3                   */
  .word _handler //     EXTI4_IRQHandler                  /* EXTI Line4                   */
  .word _handler //     DMA1_Stream0_IRQHandler           /* DMA1 Stream 0                */
  .word _handler //     DMA1_Stream1_IRQHandler           /* DMA1 Stream 1                */
  .word _handler //     DMA1_Stream2_IRQHandler           /* DMA1 Stream 2                */
  .word _handler //     DMA1_Stream3_IRQHandler           /* DMA1 Stream 3                */
  .word _handler //     DMA1_Stream4_IRQHandler           /* DMA1 Stream 4                */
  .word _handler //     DMA1_Stream5_IRQHandler           /* DMA1 Stream 5                */
  .word _handler //     DMA1_Stream6_IRQHandler           /* DMA1 Stream 6                */
  .word _handler //     ADC_IRQHandler                    /* ADC1, ADC2 and ADC3s         */
  .word _handler //     CAN1_TX_IRQHandler                /* CAN1 TX                      */
  .word _handler //     CAN1_RX0_IRQHandler               /* CAN1 RX0                     */
  .word _handler //     CAN1_RX1_IRQHandler               /* CAN1 RX1                     */
  .word _handler //     CAN1_SCE_IRQHandler               /* CAN1 SCE                     */
  .word _handler //     EXTI9_5_IRQHandler                /* External Line[9:5]s          */
  .word _handler //     TIM1_BRK_TIM9_IRQHandler          /* TIM1 Break and TIM9          */
  .word _handler //     TIM1_UP_TIM10_IRQHandler          /* TIM1 Update and TIM10        */
  .word _handler //     TIM1_TRG_COM_TIM11_IRQHandler     /* TIM1 Trigger and Commutation and TIM11 */
  .word _handler //     TIM1_CC_IRQHandler                /* TIM1 Capture Compare         */
  .word _handler //     TIM2_IRQHandler                   /* TIM2                         */
  .word _handler //     TIM3_IRQHandler                   /* TIM3                         */
  .word _handler //     TIM4_IRQHandler                   /* TIM4                         */
  .word _handler //     I2C1_EV_IRQHandler                /* I2C1 Event                   */
  .word _handler //     I2C1_ER_IRQHandler                /* I2C1 Error                   */
  .word _handler //     I2C2_EV_IRQHandler                /* I2C2 Event                   */
  .word _handler //     I2C2_ER_IRQHandler                /* I2C2 Error                   */
  .word _handler //     SPI1_IRQHandler                   /* SPI1                         */
  .word _handler //     SPI2_IRQHandler                   /* SPI2                         */
  .word _handler //     USART1_IRQHandler                 /* USART1                       */
  .word _handler //     USART2_IRQHandler                 /* USART2                       */
  .word _handler //     USART3_IRQHandler                 /* USART3                       */
  .word _handler //     EXTI15_10_IRQHandler              /* External Line[15:10]s        */
  .word _handler //     RTC_Alarm_IRQHandler              /* RTC Alarm (A and B) through EXTI Line */
  .word _handler //     OTG_FS_WKUP_IRQHandler            /* USB OTG FS Wakeup through EXTI line */
  .word _handler //     TIM8_BRK_TIM12_IRQHandler         /* TIM8 Break and TIM12         */
  .word _handler //     TIM8_UP_TIM13_IRQHandler          /* TIM8 Update and TIM13        */
  .word _handler //     TIM8_TRG_COM_TIM14_IRQHandler     /* TIM8 Trigger and Commutation and TIM14 */
  .word _handler //     TIM8_CC_IRQHandler                /* TIM8 Capture Compare         */
  .word _handler //     DMA1_Stream7_IRQHandler           /* DMA1 Stream7                 */
  .word _handler //     FMC_IRQHandler                    /* FMC                          */
  .word _handler //     SDMMC1_IRQHandler                 /* SDMMC1                       */
  .word _handler //     TIM5_IRQHandler                   /* TIM5                         */
  .word _handler //     SPI3_IRQHandler                   /* SPI3                         */
  .word _handler //     UART4_IRQHandler                  /* UART4                        */
  .word _handler //     UART5_IRQHandler                  /* UART5                        */
  .word _handler //     TIM6_DAC_IRQHandler               /* TIM6 and DAC1&2 underrun errors */
  .word _handler //     TIM7_IRQHandler                   /* TIM7                         */
  .word _handler //     DMA2_Stream0_IRQHandler           /* DMA2 Stream 0                */
  .word _handler //     DMA2_Stream1_IRQHandler           /* DMA2 Stream 1                */
  .word _handler //     DMA2_Stream2_IRQHandler           /* DMA2 Stream 2                */
  .word _handler //     DMA2_Stream3_IRQHandler           /* DMA2 Stream 3                */
  .word _handler //     DMA2_Stream4_IRQHandler           /* DMA2 Stream 4                */
  .word _handler //     ETH_IRQHandler                    /* Ethernet                     */
  .word _handler //     ETH_WKUP_IRQHandler               /* Ethernet Wakeup through EXTI line */
  .word _handler //     CAN2_TX_IRQHandler                /* CAN2 TX                      */
  .word _handler //     CAN2_RX0_IRQHandler               /* CAN2 RX0                     */
  .word _handler //     CAN2_RX1_IRQHandler               /* CAN2 RX1                     */
  .word _handler //     CAN2_SCE_IRQHandler               /* CAN2 SCE                     */
  .word _handler //     OTG_FS_IRQHandler                 /* USB OTG FS                   */
  .word _handler //     DMA2_Stream5_IRQHandler           /* DMA2 Stream 5                */
  .word _handler //     DMA2_Stream6_IRQHandler           /* DMA2 Stream 6                */
  .word _handler //     DMA2_Stream7_IRQHandler           /* DMA2 Stream 7                */
  .word _handler //     USART6_IRQHandler                 /* USART6                       */
  .word _handler //     I2C3_EV_IRQHandler                /* I2C3 event                   */
  .word _handler //     I2C3_ER_IRQHandler                /* I2C3 error                   */
  .word _handler //     OTG_HS_EP1_OUT_IRQHandler         /* USB OTG HS End Point 1 Out   */
  .word _handler //     OTG_HS_EP1_IN_IRQHandler          /* USB OTG HS End Point 1 In    */
  .word _handler //     OTG_HS_WKUP_IRQHandler            /* USB OTG HS Wakeup through EXTI */
  .word _handler //     OTG_HS_IRQHandler                 /* USB OTG HS                   */
  .word _handler //     DCMI_IRQHandler                   /* DCMI                         */
  .word _handler //     0                                 /* Reserved                     */
  .word _handler //     RNG_IRQHandler                    /* RNG                          */
  .word _handler //     FPU_IRQHandler                    /* FPU                          */
  .word _handler //     UART7_IRQHandler                  /* UART7                        */
  .word _handler //     UART8_IRQHandler                  /* UART8                        */
  .word _handler //     SPI4_IRQHandler                   /* SPI4                         */
  .word _handler //     SPI5_IRQHandler                   /* SPI5                         */
  .word _handler //     SPI6_IRQHandler                   /* SPI6                         */
  .word _handler //     SAI1_IRQHandler                   /* SAI1                         */
  .word _handler //     LTDC_IRQHandler                   /* LTDC                         */
  .word _handler //     LTDC_ER_IRQHandler                /* LTDC error                   */
  .word _handler //     DMA2D_IRQHandler                  /* DMA2D                        */
  .word _handler //     SAI2_IRQHandler                   /* SAI2                         */
  .word _handler //     QUADSPI_IRQHandler                /* QUADSPI                      */
  .word _handler //     LPTIM1_IRQHandler                 /* LPTIM1                       */
  .word _handler //     CEC_IRQHandler                    /* HDMI_CEC                     */
  .word _handler //     I2C4_EV_IRQHandler                /* I2C4 Event                   */
  .word _handler //     I2C4_ER_IRQHandler                /* I2C4 Error                   */
  .word _handler //     SPDIF_RX_IRQHandler               /* SPDIF_RX                     */
  .word _handler //     0                                 /* Reserved                     */
  .word _handler //     DFSDM1_FLT0_IRQHandler            /* DFSDM1 Filter 0 global Interrupt */
  .word _handler //     DFSDM1_FLT1_IRQHandler            /* DFSDM1 Filter 1 global Interrupt */
  .word _handler //     DFSDM1_FLT2_IRQHandler            /* DFSDM1 Filter 2 global Interrupt */
  .word _handler //     DFSDM1_FLT3_IRQHandler            /* DFSDM1 Filter 3 global Interrupt */
  .word _handler //     SDMMC2_IRQHandler                 /* SDMMC2                       */
  .word _handler //     CAN3_TX_IRQHandler                /* CAN3 TX                      */
  .word _handler //     CAN3_RX0_IRQHandler               /* CAN3 RX0                     */
  .word _handler //     CAN3_RX1_IRQHandler               /* CAN3 RX1                     */
  .word _handler //     CAN3_SCE_IRQHandler               /* CAN3 SCE                     */
  .word _handler //     JPEG_IRQHandler                   /* JPEG                         */
  .word _handler //     MDIOS_IRQHandler                  /* MDIOS                        */

# LD1 PB0 (green)
# LD2 PB7 (blue)
# LD3 PB14 (red)

# RCC 0x40023C00 - 0x40023FFF
#   RCC_AHB1ENR 0x30
# GPIOB 0x40020400
#   GPIO

.equ RCC_AHB1ENR, 0x40023C30
.equ GPIOB_MODER, 0x40020400
.equ GPIOB_ODR,   0x40020414

.global main
.section .text
.type start, %function
main:
	ldr r6, = RCC_AHB1ENR
	ldr r0, = 0x00100002 // GPIOBEN bit
	str r0, [r6]

	ldr r6, = GPIOB_MODER
	ldr r0, = 0x10000280
	str r0, [r6]

        nop
        nop
        nop

	ldr r6, = GPIOB_ODR
	mov r2, 0xffffffff
	str r2, [r6]

loop:
	nop
        nop
        nop
        nop
	b loop

_handler:
	nop
	b _handler
